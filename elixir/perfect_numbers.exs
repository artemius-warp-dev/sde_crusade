defmodule PerfectNumbers do
  @doc """
  Determine the aliquot sum of the given `number`, by summing all the factors
  of `number`, aside from `number` itself.

  Based on this sum, classify the number as:

  :perfect if the aliquot sum is equal to `number`
  :abundant if the aliquot sum is greater than `number`
  :deficient if the aliquot sum is less than `number`
  """
  @spec classify(number :: integer) :: {:ok, atom} | {:error, String.t()}
  def classify(number) do
    number
    |> get_factors()
    #|> Enum.count()
  end

  defp get_factors(number) do
    get_factors(number, [], number)
  end

  defp get_factors(_number, acc, 0), do: acc
  defp get_factors(number, acc, factor) when rem(number, factor) == 0 do
    get_factors(number, [factor | acc], factor - 1)
  end
  defp get_factors(number, acc, factor) do
    #IO.inspect(factor)
    get_factors(number, acc, factor - 1)  
  end
end

PerfectNumbers.classify(33_550_337) |> IO.inspect()
