defmodule PascalsTriangle do
  @doc """
  Calculates the rows of a pascal triangle
  with the given height
  """
  @spec rows(integer) :: [[integer]]
  def rows(num) do
    1..num
    |> Enum.map(&row/1)
  end

  # def rows(num) do
  #   Stream.iterate([1], fn row ->
      
  #     [hd(row) | Enum.map(Enum.chunk_every(row, 2, 1) |> IO.inspect(), &Enum.sum/1)]
  #   end)
  #   |> Enum.take(num)
  # end

  defp row(1), do: [1]

  defp row(n) do
    previous_line = row(n - 1) |> List.to_tuple()

    1..n
    |> Enum.map(fn x -> take(x - 1, previous_line) + take(x, previous_line) end)
  end

  defp take(i, tuple) when i == 0 or i > tuple_size(tuple), do: 0
  defp take(i, tuple), do: elem(tuple, i - 1)
end

PascalsTriangle.rows(4) |> IO.inspect()
