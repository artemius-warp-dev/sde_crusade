defmodule Minesweeper do
  @doc """
  Annotate empty spots next to mines with the number of mines next to them.
  """
  @spec annotate([String.t()]) :: [String.t()]
  def annotate([]), do: []
  def annotate([""]), do: [""]

  def annotate(board) do
    parsed_board = parse(board)
    rows = length(parsed_board)
    colums = Enum.at(parsed_board, 0) |> length()

    mines_indexes = collect_mines(parsed_board, [])

    annotate(parsed_board, rows - 1, colums - 1, 0, 0, [])
    |> merge_mines(mines_indexes)
    |> Enum.map(&Enum.join/1)
  end

  defp annotate([], _, _, _, _, _), do: []
  defp annotate([""], _, _, _, _, _), do: [""]

  defp annotate(board, rows, colums, r, c, acc) when c <= colums and r <= rows do
    score = scan(board, r, c)
    annotate(board, rows, colums, r, c + 1, [score | acc])
  end

  defp annotate(board, rows, colums, r, c, acc) when c > colums and r <= rows do
    annotate(board, rows, colums, r + 1, 0, acc)
  end

  defp annotate(_board, rows, colums, r, _c, acc) when r > rows do
    acc
    |> Enum.reverse()
    |> Enum.chunk_every(colums + 1)
  end

  defp scan(board, r, c) do
    [{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}]
    |> Enum.reduce(" ", fn {x, y}, score -> get_score(board, r + x, c + y, score) end)
  end

  defp get_score(board, r, c, score) do
    cond do
      r < 0 or c < 0 -> score
      get_in(board, [Access.at(r), Access.at(c)]) == "*" -> to_integer(score) + 1
      true -> score
    end
  end

  defp merge_mines(board, mines) do
    mines
    |> List.flatten()
    |> Enum.reduce(board, fn {r, c}, acc ->
      put_in(acc, [Access.at(r), Access.at(c)], "*")
    end)
  end

  defp collect_mines(board, acc) do
    Enum.with_index(board)
    |> Enum.map(fn {row, r} ->
      Enum.with_index(row)
      |> Enum.reduce(acc, fn {e, c}, acc ->
        if e == "*" do
          [{r, c} | acc]
        else
          acc
        end
      end)
    end)
    |> Enum.filter(fn list -> length(list) > 0 end)
  end

  defp to_integer(" "), do: 0
  defp to_integer(n), do: n

  defp parse(board) do
    board
    |> Enum.map(&String.graphemes/1)
  end
end

# input = [
#   "   ",
#   "   ",
#   " * "
# ]

input = [
  " *  * ",
  "  *   ",
  "    * ",
  "   * *",
  " *  * ",
  "      "
]

expected = [
  "1*22*1",
  "12*322",
  " 123*2",
  "112*4*",
  "1*22*2",
  "111111"
]

(Minesweeper.annotate(input) == expected)
|> IO.inspect()
