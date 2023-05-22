defmodule OcrNumbers do
  @doc """
  Given a 3 x 4 grid of pipes, underscores, and spaces, determine which number is represented, or
  whether it is garbled.
  """
  @spec convert([String.t()]) :: {:ok, String.t()} | {:error, String.t()}
  def convert(input) do
    cond do
      rem(length(input), 4) != 0 ->
        {:error, "invalid line count"}

      rem(input |> hd() |> String.length(), 3) != 0 ->
        {:error, "invalid column count"}

      true ->
        rows = Enum.chunk_every(input, 4)

        numbers_amount = input |> hd() |> String.length() |> div(3)

        rows
        |> Enum.map(fn row ->
          1..numbers_amount
          |> Enum.reduce({0, []}, fn _x, {i, acc} -> {i + 3, [parse_row(row, i) | acc]} end)
          |> then(fn {_, acc} -> acc end)
          |> Enum.reverse()
          |> Enum.map_join(&match/1)
        end)
        |> Enum.join(",")
        |> then(fn res -> {:ok, res} end)
    end
  end

  defp parse_row(input, start) do
    input
    |> Enum.reduce([], fn x, acc -> [String.slice(x, start, 3) | acc] end)
    |> Enum.reverse()
  end

  defp match([" _ ", "| |", "|_|", "   "]), do: "0"
  defp match(["   ", "  |", "  |", "   "]), do: "1"
  defp match([" _ ", " _|", "|_ ", "   "]), do: "2"
  defp match([" _ ", " _|", " _|", "   "]), do: "3"
  defp match(["   ", "|_|", "  |", "   "]), do: "4"
  defp match([" _ ", "|_ ", " _|", "   "]), do: "5"
  defp match([" _ ", "|_ ", "|_|", "   "]), do: "6"
  defp match([" _ ", "  |", "  |", "   "]), do: "7"
  defp match([" _ ", "|_|", "|_|", "   "]), do: "8"
  defp match([" _ ", "|_|", " _|", "   "]), do: "9"
  defp match(["   ", "   ", "   ", "   "]), do: ","

  defp match(_), do: "?"
end

# OcrNumbers.convert([
#   " _ ",
#   "| |",
#   "|_|",
#   "   "
# ])

OcrNumbers.convert([
  "       _     _        _  _ ",
  "  |  || |  || |  |  || || |",
  "  |  ||_|  ||_|  |  ||_||_|",
  "                           "
])
# OcrNumbers.convert([
#   "   ",
#   "  |",
#   "  |",
#   "   "
# ])

# OcrNumbers.convert([
#   "    _  _ ",
#   "  | _| _|",
#   "  ||_  _|",
#   "         ",
#   "    _  _ ",
#   "|_||_ |_ ",
#   "  | _||_|",
#   "         ",
#   " _  _  _ ",
#   "  ||_||_|",
#   "  ||_| _|",
#   "         "
# ])
|> IO.inspect()
