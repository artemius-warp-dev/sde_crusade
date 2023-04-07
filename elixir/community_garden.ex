defmodule Plot do
  @enforce_keys [:plot_id, :registered_to]
  defstruct [:plot_id, :registered_to]
end

defmodule CommunityGarden do
  def start(opts \\ []) do
    Agent.start(fn -> %{plots: [], counter: 0} end, opts)
  end

  def list_registrations(pid) do
    Agent.get(pid, fn state -> state.plots end)
  end

  def register(pid, register_to) do
    state = Agent.get(pid, fn state -> state end)
    counter = state.counter + 1
    plot = %Plot{plot_id: counter, registered_to: register_to}
    plots = [plot | state.plots]
    Agent.update(pid, fn state ->
      %{state | plots: plots, counter: counter}
    end)
    plot
  end

  def release(pid, plot_id) do
    plots=
    list_registrations(pid)
    |> Enum.filter(fn %{plot_id: id} -> id != plot_id end)
    Agent.update(pid, fn state -> %{state | plots: plots} end)
  end

  def get_registration(pid, plot_id) do
    list_registrations(pid)
    |> Enum.find(fn %{plot_id: ^plot_id} = plot -> plot end)
    |> get_registration()
  end
  defp get_registration(plot) when is_nil(plot), do: {:not_found, "plot is unregistered"}
  defp get_registration(plot), do: plot
end
