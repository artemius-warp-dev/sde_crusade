defmodule Connect do
  @doc """
  Calculates the winner (if any) of a board
  using "O" as the white player
  and "X" as the black player
  """
  @spec result_for([String.t()]) :: :none | :black | :white
  def result_for(board) do
    parsed_board = parse(board)

    transparent_board =
      parsed_board
      |> Enum.zip_reduce([], fn elements, acc -> [Enum.to_list(elements) | acc] end)
      |> Enum.map(&Enum.reverse/1)

    finish_row = parsed_board |> length()
    transparent_finish_row = transparent_board |> length()

    cond do
      find_solution_for(parsed_board, "O", 0, finish_row - 1) -> :white
      find_solution_for(transparent_board, "X", 0, transparent_finish_row - 1) -> :black
      true -> :none
    end
  end

  defp find_solution_for([first_row | _rest] = board, ch, start_row, finish_row) do
    Enum.with_index(first_row)
    |> Enum.filter(fn {v, _} -> v == ch end)
    |> Enum.map(fn {_, i} -> find_next(start_row, i, ch, board, finish_row, [{start_row, i}]) end)
    |> Enum.any?()
  end

  defp find_next(row, _column, _ch, _, row, _), do: true
  defp find_next(row, _c, _ch, _board, finish_row, _acc) when row > finish_row, do: false

  defp find_next(row, column, ch, board, finish_row, acc) do
    [{-1, 0}, {1, -1}, {1, 0}, {0, 1}, {-1, 1}, {0, -1}]
    |> Enum.any?(fn {x, y} -> check_direction(row + x, column + y, ch, board, finish_row, acc) end)
  end

  defp check_direction(row, column, ch, board, finish_row, acc)
       when row >= 0 and column >= 0 do
    elem = get_in(board, [Access.at(row), Access.at(column)])

    cond do
      elem != ch ->
        false

      Enum.find(acc, fn {r, c} -> r == row and c == column end) ->
        false

      elem == ch ->
        find_next(row, column, ch, board, finish_row, [{row, column} | acc])
    end
  end

  defp check_direction(_, _, _, _, _, _), do: false

  defp parse(board) do
    board
    |> Enum.map(&String.graphemes/1)
    |> Enum.map(&Enum.filter(&1, fn x -> x != " " end))
  end
end

# board = [
#   ". O . .",
#   " O X X X",
#   "  O X O .",
#   "   X X O X",
#   "    . O X ."
# ]

# board = [". . . . .", " . . . . .", "  . . . . .", "   . . . . .", "    . . . . ."]
board = [
  "O X X X X X X X X",
  " O X O O O O O O O",
  "  O X O X X X X X O",
  "   O X O X O O O X O",
  "    O X O X X X O X O",
  "     O X O O O X O X O",
  "      O X X X X X O X O",
  "       O O O O O O O X O",
  "        X X X X X X X X O"
]

# board = [
#   ". X X . .",
#   " X . X . X",
#   "  . X . X .",
#   "   . X X . .",
#   "    O O O O O"
# ]

# board = [
#   ". O . .",
#   " O X X X",
#   "  O O O .",
#   "   X X O X",
#   "    . O X ."
# ]

# board = ["X"]

# board = [". O . .", " O X X X", "  O X O .", "   X X O X", "    . O X ."]
# board = ["X O . .", " O X X X", "  O X O .", "   . O X .", "    X X O O"]
# board = ["X . . .", " . X O .", "  O . X O", "   . O . X", "    . . O ."
Connect.result_for(board) |> IO.inspect()
