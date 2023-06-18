defmodule React do
  @opaque cells :: pid

  @type cell :: {:input, String.t(), any} | {:output, String.t(), [String.t()], fun()}

  @doc """
  Start a reactive system
  """
  @spec new(cells :: [cell]) :: {:ok, pid}
  def new(cells) do
    ReactSrv.start_link(cells)
  end

  @doc """
  Return the value of an input or output cell
  """
  @spec get_value(cells :: pid, cell_name :: String.t()) :: any()
  def get_value(cells, cell_name) do
    GenServer.call(cells, {:get_value, cell_name})
  end

  @doc """
  Set the value of an input cell
  """
  @spec set_value(cells :: pid, cell_name :: String.t(), value :: any) :: :ok
  def set_value(cells, cell_name, value) do
    GenServer.call(cells, {:set_value, cell_name, value})
  end

  @doc """
  Add a callback to an output cell
  """
  @spec add_callback(
          cells :: pid,
          cell_name :: String.t(),
          callback_name :: String.t(),
          callback :: fun()
        ) :: :ok
  def add_callback(cells, cell_name, callback_name, callback) do
    GenServer.call(cells, {:add_callback, cell_name, callback_name, callback})
  end

  @doc """
  Remove a callback from an output cell
  """
  @spec remove_callback(cells :: pid, cell_name :: String.t(), callback_name :: String.t()) :: :ok
  def remove_callback(cells, cell_name, callback_name) do
    GenServer.call(cells, {:remove_callback, cell_name, callback_name})
  end
end

defmodule ReactSrv do
  use GenServer

  # Client

  def start_link(cells) do
    GenServer.start_link(__MODULE__, cells)
  end

  # Server (callbacks)

  @impl true
  def init(cells) do
    {:ok, cells, {:continue, cells}}
  end

  @impl true
  def handle_continue(cells, _state) do
    {:noreply, init_state(cells)}
  end

  @impl true
  def handle_call({:get_value, cell_name}, _from, state) do
    {:reply, get_value(cell_name, state), state}
  end

  @impl true
  def handle_call({:set_value, cell_name, value}, _from, state) do
    new_state = update_state(state, cell_name, value)
    {:reply, :ok, new_state}
  end

  @impl true
  def handle_call({:add_callback, cell_name, callback_name, callback}, _from, state) do
    new_state =
      add_callback(state, cell_name, callback_name, callback)

    {:reply, :ok, new_state}
  end

  @impl true
  def handle_call({:remove_callback, cell_name, callback_name}, _from, state) do
    new_state =
      remove_callback(state, cell_name, callback_name)

    {:reply, :ok, new_state}
  end

  defp remove_callback(state, cell_name, callback_name) do
    {:output, ^cell_name, input_cells, f, value, callbacks} = List.keyfind!(state, cell_name, 1)

    List.keyreplace(
      state,
      cell_name,
      1,
      {:output, cell_name, input_cells, f, value, Map.delete(callbacks, callback_name)}
    )
  end

  defp add_callback(state, cell_name, callback_name, callback) do
    {:output, ^cell_name, input_cells, f, value, callbacks} = List.keyfind!(state, cell_name, 1)

    List.keyreplace(
      state,
      cell_name,
      1,
      {:output, cell_name, input_cells, f, value, Map.put(callbacks, callback_name, callback)}
    )
  end

  defp update_state(cells, cell_name, value) do
    List.keyreplace(cells, cell_name, 1, {:input, cell_name, value})
    |> update_state()
  end

  defp update_state(cells) do
    cells
    |> Enum.reduce(cells, fn
      {:output, name, input_cells, f, old_value, callbacks}, acc ->
        new_value = calculate(input_cells, f, cells)
        do_update(input_cells, name, f, callbacks, acc, new_value, new_value != old_value)

      {_, _, _}, acc ->
        acc
    end)
  end

  defp do_update(input_cells, name, f, callbacks, acc, value, updated?) when updated? do
    Enum.map(callbacks, fn {callback_name, callback} ->
      Task.start(fn -> callback.(callback_name, value) end)
    end)

    List.keyreplace(acc, name, 1, {:output, name, input_cells, f, value, callbacks})
  end

  defp do_update(input_cells, name, f, callbacks, acc, value, _) do
    List.keyreplace(acc, name, 1, {:output, name, input_cells, f, value, callbacks})
  end

  defp init_state(cells) do
    cells
    |> Enum.reduce(cells, fn
      {:output, name, input_cells, f}, acc ->
        value = calculate(input_cells, f, cells)
        List.keyreplace(acc, name, 1, {:output, name, input_cells, f, value, %{}})

      {_, _, _}, acc ->
        acc
    end)
  end

  defp calculate(args, f, cells) do
    args
    |> Enum.map(fn name ->
      case List.keytake(cells, name, 1) do
        {{_type, ^name, value}, _} -> value
        {{_type, ^name, input_cells, f2}, _} -> calculate(input_cells, f2, cells)
        {{_type, ^name, input_cells, f3, _v3, _callbacks}, _} -> calculate(input_cells, f3, cells)
      end
    end)
    |> then(fn values -> Kernel.apply(f, values) end)
  end

  defp get_value(cell_name, state) do
    case cell_name do
      "input" ->
        {{_, ^cell_name, value}, _} = List.keytake(state, cell_name, 1)
        value

      "output" ->
        {{_, ^cell_name, _, _, value, _}, _} = List.keytake(state, cell_name, 1)
        value
    end
  end
end

send_callback = fn pid -> fn name, value -> send(pid, {:callback, name, value}) end end

# {:ok, cells} =
#  React.new([{:input, "input", 1}, {:output, "output", ["input"], fn a -> a + 1 end}])

# React.get_value(cells, "output") |> IO.inspect() == 2

# {:ok, cells} =
#   React.new([
#     {:input, "input", 1},
#     {:output, "times_two", ["input"], fn a -> a * 2 end},
#     {:output, "times_thirty", ["input"], fn a -> a * 30 end},
#     {:output, "output", ["times_two", "times_thirty"], fn a, b -> a + b end}
#   ])

# React.get_value(cells, "output") |> IO.inspect() == 32
# React.set_value(cells, "input", 3)
# React.get_value(cells, "output") |> IO.inspect() == 96

# {:ok, cells} =
#   React.new([
#     {:input, "one", 1},
#     {:input, "two", 2},
#     {:output, "output", ["one", "two"], fn a, b -> a + b * 10 end}
#   ])

# React.get_value(cells, "output") |> IO.inspect() == 21

# {:ok, cells} =
#   React.new([
#     {:input, "input", 1},
#     {:output, "output", ["input"], fn a -> a + 1 end}
#   ])

# myself = self()
# React.add_callback(cells, "output", "callback1", send_callback.(myself))
# React.set_value(cells, "input", 3)

# Some incorrect implementations store their callbacks in an array

# and removing a callback repeatedly either removes an unrelated callback

# or causes an out of bounds access.

# {:ok, cells} =
#   React.new([
#     {:input, "input", 1},
#     {:output, "output", ["input"], fn a -> a + 1 end}
#   ])

# myself = self()
# React.add_callback(cells, "output", "callback1", send_callback.(myself))
# React.add_callback(cells, "output", "callback2", send_callback.(myself))
# React.remove_callback(cells, "output", "callback1")
# React.remove_callback(cells, "output", "callback1")
# React.remove_callback(cells, "output", "callback1")
# React.set_value(cells, "input", 2)

# refute_receive {:callback, "callback1", _}, 50, "Expected no callback"

# assert_receive {:callback, "callback2", 3}

# receive do
#  {:callback, "callback2", 3} -> IO.inspect("HELLO")
# end

# assert_receive {:callback, "callback1", 4}

# # Some incorrect implementations call a callback function too early,
# # when not all of the inputs of a compute cell have propagated new values.
# {:ok, cells} =
#   React.new([
#     {:input, "input", 1},
#     {:output, "plus_one", ["input"], fn a -> a + 1 end},
#     {:output, "minus_one1", ["input"], fn a -> a - 1 end},
#     {:output, "minus_one2", ["minus_one1"], fn a -> a - 1 end},
#     {:output, "output", ["plus_one", "minus_one2"], fn a, b -> a * b end}
#   ])

# myself = self()
# React.add_callback(cells, "output", "callback1", send_callback(myself))
# React.set_value(cells, "input", 4)
# assert_receive {:callback, "callback1", 10}


{:ok, cells} =
  React.new([
    {:input, "input", 1},
    {:output, "output", ["input"],
     fn a ->
       if a < 3 do
         111
       else
         222
       end
     end}
  ])

myself = self()
React.add_callback(cells, "output", "callback1", send_callback.(myself))
React.set_value(cells, "input", 2)
#refute_receive {:callback, _, _}, 50, "Expected no callback"
React.set_value(cells, "input", 4)
#assert_receive {:callback, "callback1", 222}
