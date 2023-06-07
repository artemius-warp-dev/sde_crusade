defmodule Spiral do
  @doc """
  Given the dimension, return a square matrix of numbers in clockwise spiral order.
  """
  @spec matrix(dimension :: integer) :: list(list(integer))
  def matrix(0), do: []
  def matrix(dimension) do
    matrix = prepare_matrix(dimension)
    do_spiral({matrix, 0}, 0, 0, dimension - 1, dimension - 1)
  end

  defp do_spiral({matrix, n}, top, left, bottom, right) do
    if top > bottom or left > right do
      matrix
    else
      {matrix, n}
      |> left_to_right(top, top, right)
      |> top_from_bottom(right, top + 1, bottom)
      |> right_to_left(bottom, right - 1, left)
      |> bottom_to_top(left, bottom - 1, top + 1)
      |> do_spiral(top + 1, left + 1, bottom - 1, right - 1)
    end
  end

  defp bottom_to_top({matrix, n}, left, i, top) when i >= top do
    updated_matrix = put_in(matrix, [Access.at(i), Access.at(left)], n + 1)

    {updated_matrix, n + 1}
    |> bottom_to_top(left, i - 1, top)
  end

  defp bottom_to_top({matrix, n}, _, _, _), do: {matrix, n}

  defp right_to_left({matrix, n}, bottom, i, left) when i >= left do
    updated_matrix = put_in(matrix, [Access.at(bottom), Access.at(i)], n + 1)

    {updated_matrix, n + 1}
    |> right_to_left(bottom, i - 1, left)
  end

  defp right_to_left({matrix, n}, _, _, _), do: {matrix, n}

  defp left_to_right({matrix, n}, top, i, right) when i <= right do
    updated_matrix = put_in(matrix, [Access.at(top), Access.at(i)], n + 1)

    {updated_matrix, n + 1}
    |> left_to_right(top, i + 1, right)
  end

  defp left_to_right({matrix, n}, _, _, _), do: {matrix, n}

  defp top_from_bottom({matrix, n}, right, i, bottom) when i <= bottom do
    updated_matrix = put_in(matrix, [Access.at(i), Access.at(right)], n + 1)

    {updated_matrix, n + 1}
    |> top_from_bottom(right, i + 1, bottom)
  end

  defp top_from_bottom({matrix, n}, _, _, _), do: {matrix, n}

  defp fill(n, size) do
    Stream.iterate(n, &(&1 + 1)) |> Enum.take(size)
  end

  defp prepare_matrix(number) do
    for _x <- 1..number, do: for(_y <- 1..number, do: -1)
  end
end

# defmodule Spiral do
#   @doc """
#   Given the dimension, return a square matrix of numbers in clockwise spiral order.
#   Inspired by dantetekanem's solution at
#   https://exercism.io/tracks/elixir/exercises/spiral-matrix/solutions/052c481bebf845a9992d88d94a6aae8b
#   (mainly, improve the naming)
#   """

#   @spec matrix(dimension :: integer) :: list(list(integer))

#   def matrix(0), do: []

#   def matrix(dimension), do: do_matrix(dimension, dimension, 1)

#   defp do_matrix(_, 0, _), do: [[]]

#   defp do_matrix(rows, cols, min) do
#     top_row = Enum.to_list(min..(min + cols - 1)) |> IO.inspect(charlists: :as_lists, label: "FIRST_ROW")
    
#     other_rows = do_matrix(cols, rows - 1, min + cols) |> rotate_right |> IO.inspect(label: "OTHER_ROWS")

#     [top_row | other_rows]
#   end

#   defp rotate_right(rows), do: rows |> Enum.reverse() |> transpose

#   defp transpose(rows), do: rows |> IO.inspect(charlists: :as_lists, label: "BEFORE TRANSPOSE") |> Enum.zip() |> Enum.map(&Tuple.to_list/1)
# end

Spiral.matrix(3) |> IO.inspect()
