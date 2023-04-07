defmodule TakeANumber do
  def start() do
    spawn(fn ->
      receive_loop(0)
    end)
  end

  defp receive_loop(state) do
    receive do
      {:report_state, sender_pid} ->
        send(sender_pid, state)
        receive_loop(state)

      {:take_a_number, sender_pid} ->
        new_state = state + 1
        send(sender_pid, new_state)
        receive_loop(new_state)

      {:stop} ->
        :stopped

      _ ->
        receive_loop(state)
    end
  end
end
