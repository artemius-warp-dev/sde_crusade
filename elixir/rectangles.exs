defmodule Rectangles do
  @doc """
  Count the number of ASCII rectangles.
  """
  @spec count(input :: String.t()) :: integer
  def count(input) do
    parsed = parse(input)
    corners = find_corners(parsed) |> List.flatten()

    for {ar, ac} = a <- corners,
        {br, bc} = b <- corners,
        {cr, cc} = c <- corners,
        {dr, dc} = d <- corners,
        ac == cc and bc == dc and ar == br and cr == dr and ar != cr and ac != bc and cc != dc,
        full_row?(parsed, a, b) and full_row?(parsed, c, d),
        full_column?(parsed, a, c) and full_column?(parsed, b, d) do
      [a, b, c, d] |> Enum.sort()
    end
    |> Enum.uniq()
    |> Enum.count()
  end

  defp full_row?(matrix, {r, c1}, {r, c2}) do
    c1..c2
    |> Enum.map(&point?(matrix, r, &1, "-"))
    |> Enum.all?()
  end

  defp full_column?(matrix, {r1, c}, {r2, c}) do
    r1..r2
    |> Enum.map(&point?(matrix, &1, c, "|"))
    |> Enum.all?()
  end

  defp point?(matrix, r, c, ch) do
    {columns, _row} = get_in(matrix, [Access.at(r)])
    {elem, _} = get_in(columns, [Access.at(c)])
    elem == ch or elem == "+"
  end

  defp find_corners(input) do
    input
    |> Enum.reduce([], fn row, acc ->
      [find_x_positions(row) | acc]
    end)
    |> Enum.reverse()
  end

  defp find_x_positions({row, r}) do
    row
    |> Enum.reduce([], fn
      {"+", i}, acc ->
        [{r, i} | acc]

      _, acc ->
        acc
    end)
    |> Enum.reverse()
  end

  defp parse(input) do
    input
    |> String.split("\n")
    |> Enum.map(&(String.graphemes(&1) |> Enum.with_index()))
    |> Enum.filter(fn x -> x != [] end)
    |> Enum.with_index()
  end
end

# input = """
# +-+
# | |
# +-+
# """

# input = """
#   +-+
#   | |
# +-+-+
# | | \s
# +-+ \s
# """
input = """
   +--+
  ++  |
+-++--+
|  |  |
+--+--+
"""

# input = """
#   +-+
#     |
# +-+-+
# | | -
# +-+-+
# """

Rectangles.count(input) |> IO.inspect()
