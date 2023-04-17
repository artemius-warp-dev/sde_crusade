defmodule SumOfMultiples do
  @doc """
  Adds up all numbers from 1 to a given end number that are multiples of the factors provided.
  """
  @spec to(non_neg_integer, [non_neg_integer]) :: non_neg_integer
  def to(limit, factors) do
    factors
    |> Enum.map(&(find_multiples(&1, limit)))
    |> Enum.reduce([], &concat/2)
    |> Enum.sum()
  end

  defp find_multiples(0, _), do: []
  defp find_multiples(factor, limit) do
    1..limit-1
    |> Enum.filter(&(rem(&1,factor) == 0))
  end

  defp concat(list, acc) do
    Enum.concat(list, acc)
    |> Enum.uniq()
  end
end
