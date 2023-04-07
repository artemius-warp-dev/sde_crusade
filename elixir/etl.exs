defmodule ETL do
  @doc """
  Transforms an old Scrabble score system to a new one.

  ## Examples

    iex> ETL.transform(%{1 => ["A", "E"], 2 => ["D", "G"]})
    %{"a" => 1, "d" => 2, "e" => 1, "g" => 2}
  """
  @spec transform(map) :: map
  def transform(input) do
    for {key, values} <- input,
        value <- values,
        into: %{} do
      {String.downcase(value), key}
    end
  end

  # def transform(input) do
  #   input
  #   |> Enum.reduce(%{}, fn m, acc ->
  #     {key, values} = m
  #     map = values |> new_map(key)
  #     Map.merge(acc, map)
  #   end)
  # end

  # defp new_map(keys, value) do
  #   keys
  #   |> Enum.map(&%{String.downcase(&1) => value})
  #   |> Enum.reduce(%{}, &Map.merge/2)
  # end
end

ETL.transform(%{1 => ["A", "E"], 2 => ["D", "G"]}) |> IO.inspect()
