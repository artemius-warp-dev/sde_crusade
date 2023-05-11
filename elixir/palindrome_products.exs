defmodule PalindromeProducts do
  @doc """
  Generates all palindrome products from an optionally given min factor (or 1) to a given max factor.
  """
  @spec generate(non_neg_integer, non_neg_integer) :: map
  def generate(max_factor, min_factor \\ 1)
  def generate(max_factor, min_factor) when max_factor < min_factor, do: raise(ArgumentError)

  def generate(max_factor, min_factor) do
    palindrom_factors =
      for x <- min_factor..max_factor,
          y <- x..max_factor,
          factor = x * y,
          palindrom?(factor) do
        {factor, [x, y]}
      end

    palindrom_factors
    |> Enum.reduce(%{}, fn {k, [x, y]}, acc ->
      Map.update(acc, k, [[x, y]], fn values ->
        cond do
          [x, y] not in values and [y, x] not in values ->
            [[x, y] | values] |> Enum.sort_by(fn [a, _b] = enums -> {a, enums} end, :asc)

          true ->
            values
        end
      end)
    end)
  end

  defp palindrom?(number) do
    number == Integer.undigits(Enum.reverse(Integer.digits(number)))
  end
end

PalindromeProducts.generate(9999, 1000) |> IO.inspect()
