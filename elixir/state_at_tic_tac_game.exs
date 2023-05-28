defmodule StateOfTicTacToe do
  @doc """
  Determine the state a game of tic-tac-toe where X starts.
  """
  @spec game_state(board :: String.t()) :: {:ok, :win | :ongoing | :draw} | {:error, String.t()}
  def game_state(board) do
    solutions = [win_in_row?(board), win_in_column?(board), win_in_diagonal?(board)]

    cond do
      x_went_twice?(board) ->
        {:error, "Wrong turn order: X went twice"}

      o_started?(board) ->
        {:error, "Wrong turn order: O started"}

      solutions |> Enum.sum() == 1 ->
        {:ok, :win}

      solutions == [1, 1, 0] ->
        {:ok, :win}

      solutions == [0, 0, 2] ->
        {:ok, :win}

      solutions |> Enum.sum() > 1 ->
        {:error, "Impossible board: game should have ended after the game was won"}

      ongoing?(board) ->
        {:ok, :ongoing}

      true ->
        {:ok, :draw}
    end
  end

  defp match?(row) do
    element = List.first(row)
    Enum.all?(row, fn x -> element == x and x != "." end)
  end

  defp parse(board) do
    board
    |> String.graphemes()
    |> Enum.filter(fn x -> x in ["X", "O", "."] end)
    |> Enum.chunk_every(3)
  end

  defp ongoing?(board) do
    parse(board)
    |> Enum.any?(&Enum.any?(&1, fn x -> x == "." end))
  end

  defp win_in_row?(board) do
    parse(board)
    |> Enum.map(&match?/1)
    |> Enum.count(fn res -> res == true end)
  end

  defp win_in_column?(board) do
    parse(board)
    |> Enum.zip_reduce([], fn elements, acc -> [elements | acc] end)
    |> Enum.map(&match?/1)
    |> Enum.count(fn res -> res == true end)
  end

  defp win_in_diagonal?(board) do
    rows = parse(board)
    left_diagonal = get_diagonal(rows, 0, 1)
    right_diagonal = get_diagonal(rows, 2, -1)

    [match?(left_diagonal), match?(right_diagonal)]
    |> Enum.count(fn res -> res == true end)
  end

  defp get_diagonal(rows, start, factor) do
    rows
    |> Enum.reduce({[], start}, fn row, {acc, i} -> {[Enum.at(row, i) | acc], i + factor} end)
    |> then(fn {diagonal, _} -> diagonal end)
  end

  defp calculate_amounts(board) do
    list = List.flatten(parse(board))
    amount_of_x = list |> Enum.count(fn x -> x == "X" end)
    amount_of_o = list |> Enum.count(fn x -> x == "O" end)
    {amount_of_x, amount_of_o}
  end

  defp o_started?(board) do
    {amount_of_x, amount_of_o} = calculate_amounts(board)
    amount_of_o - amount_of_x >= 1
  end

  defp x_went_twice?(board) do
    {amount_of_x, amount_of_o} = calculate_amounts(board)
    amount_of_x - amount_of_o >= 2
  end
end

board = """
XXX
XOO
XOO
"""

# [1 2 3]
# [3 2 1]

StateOfTicTacToe.game_state(board) |> IO.inspect()
