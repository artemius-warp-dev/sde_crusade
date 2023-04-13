defmodule StringSeries do
  @doc """
  Given a string `s` and a positive integer `size`, return all substrings
  of that size. If `size` is greater than the length of `s`, or less than 1,
  return an empty list.
  """
  @spec slices(s :: String.t(), size :: integer) :: list(String.t())
  def slices(s, 1) do
    String.codepoints(s)
    |> Enum.chunk_every(1)
    |> Enum.map(&Enum.join/1)
  end

  def slices(_s, size) when size <= 0, do: []

  def slices(s, size) do
    len = String.length(s)
    slice(s, size, len)
  end

  defp slice(_s, size, length) when size > length, do: []

  defp slice(s, size, length) do
    for x <- 0..(length - size), do: String.slice(s, x, size)
  end
end
