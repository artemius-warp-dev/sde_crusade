defmodule Queens do
  @type t :: %Queens{black: {integer, integer}, white: {integer, integer}}
  defstruct [:white, :black]

  @doc """
  Creates a new set of Queens
  """
  @spec new(Keyword.t()) :: Queens.t()
  def new(opts \\ [])
  def new(white: {r, c}, black: {r, c}), do: raise(ArgumentError, "invalid argument")
  def new(black: {r, c}, white: {r, c}), do: raise(ArgumentError, "invalid argument")

  def new(opts) do
    Keyword.keys(opts)
    |> Enum.reduce(%Queens{}, fn key, acc ->
      case Keyword.get(opts, key) do
        nil ->
          acc

        {r, c} when r >= 0 and r < 8 and c >= 0 and c < 8 and key in [:white, :black] ->
          Map.put(acc, key, {r, c})

        _ ->
          raise(ArgumentError, "invalid argument")
      end
    end)
  end

  @doc """
  Gives a string representation of the board with
  white and black queen locations shown
  """
  @spec to_string(Queens.t()) :: String.t()
  def to_string(queens) do
    queens
    |> create_board()
    |> Enum.map(&Enum.join(&1, " "))
    |> Enum.join("\n")
  end

  @doc """
  Checks if the queens can attack each other
  """
  @spec can_attack?(Queens.t()) :: boolean
  def can_attack?(%Queens{white: {wr, wc}, black: _} = queens) do
    board = create_board(queens)

    [
      attack_at_rows?(board, 0, wc, "B"),
      attack_at_colums?(board, wr, 0, "B"),
      attack_at_diagonals?(board, wr, wc, "B", [])
    ]
    |> Enum.any?()
  end

  def can_attack?(_), do: false

  defp create_board(queens) do
    board = for _x <- 1..8, do: for(_y <- 1..8, do: "_")

    Map.from_struct(queens)
    |> Enum.reduce(board, fn
      {key, {r, c}}, acc ->
        value =
          case key do
            :white -> "W"
            :black -> "B"
          end

        put_in(acc, [Access.at(r), Access.at(c)], value)

      _, acc ->
        acc
    end)
  end

  defp attack_at_rows?(board, r, c, queen) when r < 8 do
    if get_in(board, [Access.at(r), Access.at(c)]) == queen do
      true
    else
      attack_at_rows?(board, r + 1, c, queen)
    end
  end

  defp attack_at_rows?(_, _, _, _), do: false

  defp attack_at_colums?(board, r, c, queen) when c < 8 do
    if get_in(board, [Access.at(r), Access.at(c)]) == queen do
      true
    else
      attack_at_colums?(board, r, c + 1, queen)
    end
  end

  defp attack_at_colums?(_, _, _, _), do: false

  defp attack_at_diagonals?(board, r, c, queen, acc) do
    attack_at_left_diagonal?(board, r, c, queen, acc) or
      attack_at_right_diagonal?(board, r, c, queen, acc)
  end

  defp queen_exist?(board, r, c, queen) do
    get_in(board, [Access.at(r), Access.at(c)]) == queen
  end

  defp attack_at_left_diagonal?(board, r, c, queen, acc)
       when r < 8 and r >= 0 and c < 8 and c >= 0 do
    visited? =
      Enum.any?(acc, fn
        {ar, ac} ->
          ar == r and ac == c

        _acc ->
          nil
      end)

    if queen_exist?(board, r, c, queen) do
      true
    else
      not visited? and
        (attack_at_left_diagonal?(board, r - 1, c + 1, queen, [{r, c} | acc]) or
           attack_at_left_diagonal?(board, r + 1, c - 1, queen, [{r, c} | acc]))
    end
  end

  defp attack_at_left_diagonal?(_, _, _, _, _), do: false

  defp attack_at_right_diagonal?(board, r, c, queen, acc)
       when r < 8 and r >= 0 and c < 8 and c >= 0 do
    visited? =
      Enum.any?(acc, fn
        {ar, ac} ->
          ar == r and ac == c

        _acc ->
          nil
      end)

    if queen_exist?(board, r, c, queen) do
      true
    else
      not visited? and
        (attack_at_right_diagonal?(board, r + 1, c + 1, queen, [{r, c} | acc]) or
           attack_at_right_diagonal?(board, r - 1, c - 1, queen, [{r, c} | acc]))
    end
  end

  defp attack_at_right_diagonal?(_, _, _, _, _), do: false
end

# queens = Queens.new(white: {1, 7}, black: {0, 6}) |> IO.inspect()
# queens = Queens.new(white: {5, 4}, black: {2, 4})
# queens = Queens.new(white: {2, 2}, black: {0, 4})
# queens = Queens.new(white: {2, 4}, black: {6, 6})
# Queens.can_attack?(queens) |> IO.inspect()
queens = Queens.new(white: {2, 4})
Queens.to_string(queens) |> IO.inspect()
