defmodule SquareRoot do
  @doc """
  Calculate the integer square root of a positive integer
  """
  @spec calculate(radicand :: pos_integer) :: pos_integer
  def calculate(radicand) do
    1..radicand
    |> Enum.find(fn x -> x * x == radicand end)
  end
end



SquareRoot.calculate(65025) |> IO.inspect()
