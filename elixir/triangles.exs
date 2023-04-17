defmodule Triangle do
  @type kind :: :equilateral | :isosceles | :scalene

  @doc """
  Return the kind of triangle of a triangle with 'a', 'b' and 'c' as lengths.
  """
  @spec kind(number, number, number) :: {:ok, kind} | {:error, String.t()}
  def kind(a, b, c) do
    cond do
      a <= 0 or b <= 0 or c <= 0 -> {:error, "all side lengths must be positive"}
      not triangle?(a,b,c) -> {:error, "side lengths violate triangle inequality"}
      a == b and a == c -> {:ok, :equilateral}
      a == b or c == b or c == a -> {:ok, :isosceles}
      not (a == b == c) -> {:ok, :scalene}
    end
  end

  defp triangle?(a,b,c) when a > 0 and b > 0 and c > 0 do
    a + b >= c and
    b + c >= a and
    a + c >= b
  end
   
end
