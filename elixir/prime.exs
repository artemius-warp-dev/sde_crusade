defmodule Prime do
  @doc """
  Generates the nth prime.
  """
  @spec nth(non_neg_integer) :: non_neg_integer
  def nth(count) when count > 0 do
    Stream.iterate(2, &(&1 + 1))
    |> Stream.filter(&prime?/1)
    |> Enum.take(count)
    |> List.last()

    # do_count(count, 1)
  end

  defp do_count(count, n), do: do_count(count, n, n)
  defp do_count(0, _next, prev), do: prev

  defp do_count(count, next, _prev) do
    prime?(next)
    |> get_count(count)
    |> do_count(next + 1, next)
  end

  defp get_count(prime?, count) when prime?, do: count - 1
  defp get_count(_, count), do: count

  defp prime?(1), do: false
  defp prime?(2), do: true
  defp prime?(number), do: prime?(number, 2)

  defp prime?(number, number), do: true
  defp prime?(number, multiplier) when rem(number, multiplier) == 0, do: false
  defp prime?(number, multiplier), do: prime?(number, multiplier + 1)

  # defp prime?(number) do
  #   1..number
  #   |> Enum.filter(&(rem(number, &1) == 0))
  #   |> Kernel.==([1, number])
  # end
end

Prime.nth(6) |> IO.inspect()
