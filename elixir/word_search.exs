defmodule WordSearch do
  defmodule Location do
    defstruct [:from, :to]

    @type t :: %Location{
            from: %{row: integer, column: integer},
            to: %{row: integer, column: integer}
          }
  end

  @doc """
  Find the start and end positions of words in a grid of letters.
  Row and column positions are 1 indexed.
  """
  @spec search(grid :: String.t(), words :: [String.t()]) :: %{String.t() => nil | Location.t()}
  def search(grid, words) do
    matrix = parse_to_matrix(grid)
    mrows = length(matrix) - 1
    mcolumns = (matrix |> hd |> length()) - 1

    words
    |> Enum.map(&search_by_words(&1, matrix, mrows, mcolumns))
    |> Map.new()
  end

  defp search_by_words(word, matrix, mrows, mcolumns) do
    word
    |> String.first()
    |> find_start_points(matrix)
    |> Enum.map(fn {r, c} -> search_by_directions(word, matrix, mrows, mcolumns, r, c, []) end)
    |> List.flatten()
    |> then(fn
      res when res == [] ->
        {word, nil}

      res ->
        start_ch = String.first(word)
        end_ch = String.last(word)

        {_start_ch, {from_r, from_c}} = Enum.find(res, fn {xch, _coords} -> start_ch == xch end)
        {_end_ch, {to_r, to_c}} = Enum.find(res, fn {xch, _coords} -> end_ch == xch end)

        {word,
         %Location{
           from: %{row: from_r + 1, column: from_c + 1},
           to: %{row: to_r + 1, column: to_c + 1}
         }}
    end)
  end

  defp search_by_directions(word, matrix, mrows, mcolumns, r, c, acc) do
    [{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}]
    |> Enum.map(&search_by_characters(word, matrix, mrows, mcolumns, r, c, acc, &1))
  end

  defp search_by_characters(
         <<ch::binary-size(1), rest::binary>>,
         matrix,
         mrows,
         mcolumns,
         r,
         c,
         acc,
         {dx, dy}
       ) do
    if point?(ch, matrix, mrows, mcolumns, r, c) do
      search_by_characters(
        rest,
        matrix,
        mrows,
        mcolumns,
        r + dx,
        c + dy,
        [{ch, {r, c}} | acc],
        {dx, dy}
      )
    else
      []
    end
  end

  defp search_by_characters("", _matrix, _mrows, _mcolumns, _r, _c, acc, _), do: acc

  defp find_start_points(ch, matrix) do
    for {row, r} <- Enum.with_index(matrix),
        {element, c} <- Enum.with_index(row),
        element == ch do
      {r, c}
    end
  end

  defp point?(ch, matrix, mrows, mcolumns, r, c)
       when r >= 0 and c >= 0 and r <= mrows and c <= mcolumns do
    get_in(matrix, [Access.at(r), Access.at(c)]) == ch
  end

  defp point?(_ch, _matr, _mrows, _mcolumns, _r, _c), do: false

  defp parse_to_matrix(grid) do
    grid
    |> String.split("\n")
    |> Enum.map(&String.graphemes/1)
    |> Enum.reject(fn x -> x == [] end)
  end
end

# grid = "clojurermt"
# grid = "mtclojurer"
# words = ["clojure"]

# grid = "xcoffeezlp"
# words = ["coffee"]

# grid = """
# jefblpepre
# tclojurerm
# """

# grid = """
# camdcimgtc
# jefblpepre
# clojurermt
# """

# grid = """
# jefblpepre
# camdcimgtc
# oivokprjsm
# pbwasqroua
# rixilelhrs
# wolcqlirpc
# screeaumgr
# alxhpburyi
# jalaycalmp
# clojurermt
# """

# words = ["clojure"]

# grid = """
# jefblpepre
# camdcimgtc
# oivokprjsm
# pbwasqroua
# rixilelhrs
# wolcqlirpc
# fortranftw
# alxhpburyi
# jalaycalmp
# clojurermt
# """

# words = ["fortran", "clojure"]

grid = """
jefblpepre
camdcimgtc
oivokprjsm
pbwasqroua
rixilelhrs
wolcqlirpc
screeaumgr
alxhpburyi
jalaycalmp
clojurermt
"""

words = [
  # "clojure",
  # "elixir",
  # "ecmascript",
  "rust",
  "java",
  "lua",
  "lisp"
  # "ruby",
  # "haskell"
]

output = WordSearch.search(grid, words) |> IO.inspect()

# expected = %{
#   "clojure" => %Location{from: %{row: 1, column: 1}, to: %{row: 1, column: 7}}
# }
