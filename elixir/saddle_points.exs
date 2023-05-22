defmodule SaddlePoints do
  @doc """
  Parses a string representation of a matrix
  to a list of rows
  """
  @spec rows(String.t()) :: [[integer]]
  def rows(""), do: []

  def rows(str) do
    str
    |> String.split("\n")
    |> Enum.map(&normalize/1)
  end

  defp normalize(str) do
    str
    |> String.split()
    |> Enum.map(&String.to_integer/1)
  end

  @doc """
  Parses a string representation of a matrix
  to a list of columns
  """
  @spec columns(String.t()) :: [[integer]]
  def columns(str) do
    rows(str)
    |> Enum.zip()
    |> Enum.map(&Tuple.to_list/1)
  end

  @doc """
  Calculates all the saddle points from a string
  representation of a matrix
  """
  @spec saddle_points(String.t()) :: [{integer, integer}]
  def saddle_points(str) do
    largest = Enum.map(rows(str), &Enum.max/1) |> Enum.with_index(1)
    smallest = Enum.map(columns(str), &Enum.min/1) |> Enum.with_index(1)

    Enum.reduce(largest, [], fn coord, acc -> match_point(coord, smallest, acc) end)
    |> Enum.reverse()
  end

  defp match_point({v, i}, smallest, acc) do
    smallest
    |> Enum.filter(fn {k, _i} -> k == v end)
    |> Enum.reduce(acc, fn {_k, j}, list -> [{i, j} | list] end)
  end
end
