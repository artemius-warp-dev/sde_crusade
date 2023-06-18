defmodule CircularBuffer do
  @moduledoc """
  An API to a stateful process that fills and empties a circular buffer
  """

  @doc """
  Create a new buffer of a given capacity
  """
  @spec new(capacity :: integer) :: {:ok, pid}
  def new(capacity) do
    BufferServer.start_link(capacity)
  end

  @doc """
  Read the oldest entry in the buffer, fail if it is empty
  """
  @spec read(buffer :: pid) :: {:ok, any} | {:error, atom}
  def read(buffer) do
    GenServer.call(buffer, :read)
  end

  @doc """
  Write a new item in the buffer, fail if is full
  """
  @spec write(buffer :: pid, item :: any) :: :ok | {:error, atom}
  def write(buffer, item) do
    GenServer.call(buffer, {:write, item})
  end

  @doc """
  Write an item in the buffer, overwrite the oldest entry if it is full
  """
  @spec overwrite(buffer :: pid, item :: any) :: :ok
  def overwrite(buffer, item) do
    GenServer.call(buffer, {:overwrite, item})
  end

  @doc """
  Clear the buffer
  """
  @spec clear(buffer :: pid) :: :ok
  def clear(buffer) do
    GenServer.call(buffer, :clear)
  end
end

defmodule BufferServer do
  use GenServer

  # Client

  def start_link(capacity) do
    GenServer.start_link(__MODULE__, capacity)
  end

  # Server (callbacks)

  @impl true
  def init(capacity) do
    {:ok, new(capacity)}
  end

  @impl true
  def handle_call(:clear, _from, state) do
    {:reply, :ok, new(state.capacity)}
  end

  @impl true
  def handle_call({:write, element}, _from, state) do
    {res, updated_state} =
      if free?(state) do
        {:ok, %{state | buffer: [element | state.buffer]} |> update_head()}
      else
        {{:error, :full}, state}
      end

    {:reply, res, updated_state}
  end

  @impl true
  def handle_call({:overwrite, element}, _from, state) do
    updated_state =
      if free?(state) do
        %{state | buffer: [element | state.buffer]} |> update_head()
      else
        %{state | buffer: [element | state.buffer], head: nil}
        |> update_head()
      end


    {:reply, :ok, updated_state}
  end

  @impl true
  def handle_call(:read, _from, state) do
    {res, updated_state} =
      if not empty?(state) do
        {{:ok, state.head}, %{state | head: nil} |> update_head()}
      else
        {{:error, :empty}, state}
      end

    {:reply, res, updated_state}
  end

  defp free?(state) do
    length(state.buffer) + head(state.head) < state.capacity
  end

  defp empty?(state) do
    is_nil(state.head)
  end

  defp head(nil), do: 0
  defp head(_n), do: 1

  defp update_head(%{head: nil, buffer: buffer} = state) when length(buffer) > 0 do
    buffer
    |> Enum.reverse()
    |> then(fn [h | rest] -> %{head: h, buffer: Enum.reverse(rest)} end)
    |> then(fn updated_state -> Map.merge(state, updated_state) end)
  end

  defp update_head(state), do: state

  defp new(capacity) do
    %{buffer: [], capacity: capacity, head: nil}
  end
end

# capacity = 2

# {:ok, buffer} = CircularBuffer.new(capacity)
# # |> IO.inspect()
# CircularBuffer.clear(buffer)
# # # :ok
# CircularBuffer.write(buffer, 1) |> IO.inspect()
# # # :ok
# CircularBuffer.write(buffer, 2) |> IO.inspect()
# CircularBuffer.overwrite(buffer, 3) |> IO.inspect()
# CircularBuffer.overwrite(buffer, 4) |> IO.inspect()
# # # {:ok, 3}
# CircularBuffer.read(buffer) |> IO.inspect()
# # # {:ok, 4}
# CircularBuffer.read(buffer) |> IO.inspect()
# # # {:error, :empty}
# CircularBuffer.read(buffer) |> IO.inspect()

# capacity = 1
# {:ok, buffer} = CircularBuffer.new(capacity)
# CircularBuffer.write(buffer, 1) |> IO.inspect() #:ok
# CircularBuffer.write(buffer, 2) |> IO.inspect() # {:error, :full}

capacity = 2
{:ok, buffer} = CircularBuffer.new(capacity)
# == :ok
CircularBuffer.write(buffer, 1)
CircularBuffer.overwrite(buffer, 2)
# == {:ok, 1}
CircularBuffer.read(buffer) |> IO.inspect()
# == {:ok, 2}
CircularBuffer.read(buffer) |> IO.inspect()
