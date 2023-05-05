defmodule Series do
  @doc """
  Finds the largest product of a given number of consecutive numbers in a given string of numbers.
  """
  @spec largest_product(String.t(), non_neg_integer) :: non_neg_integer
  def largest_product(number_string, size) do
    number_string
    |> String.graphemes()
    |> Enum.chunk_every(size, 1)
    # |> IO.inspect()
    |> Enum.map(&sum/1)
    |> Enum.sort(:desc)
    # |> IO.inspect()
    |> List.first()
  end

  defp sum(str_list) do
    str_list
    |> Enum.map(&String.to_integer/1)
    |> IO.inspect()
    |> Enum.reduce(1, fn x, acc -> x * acc end)
  end
end

Series.largest_product("63915", 3) |> IO.inspect()
