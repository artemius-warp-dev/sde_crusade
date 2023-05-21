defmodule Sieve do
  @doc """
  Generates a list of primes up to a given limit.
  """
  @spec primes_to(non_neg_integer) :: [non_neg_integer]
  def primes_to(1), do: []

  def primes_to(limit) do
    for x <- 2..limit do
      {x, :unmarked}
    end
    |> Map.new()
    |> mark_numbers(2, 2, limit)
    |> Enum.filter(fn {k, v} -> v == :unmarked end)
    |> Map.new()
    |> Map.keys()
    |> Enum.sort()
  end


  defp mark_numbers(map, factor, number, limit) when number > limit do
    case get_new_factor(map, factor + 1) do
      :finish -> map
      new_factor -> mark_numbers(map, new_factor, new_factor, limit)
    end
  end

  defp mark_numbers(map, factor, number, limit) do
    Map.replace(map, number + factor, :marked)
    |> mark_numbers(factor, number + factor, limit)
  end

  defp get_new_factor(map, factor) do
    case Map.get(map, factor) do
      nil -> :finish
      :marked -> get_new_factor(map, factor + 1)
      :unmarked -> factor
    end
  end
end

Sieve.primes_to(1000) |> IO.inspect()
