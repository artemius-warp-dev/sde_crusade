defmodule Isogram do
  @doc """
  Determines if a word or sentence is an isogram
  """
  @spec isogram?(String.t()) :: boolean
  def isogram?(sentence) do
    
    Regex.scan(~r/[[:alpha:]]/, String.downcase(sentence))
    |> List.flatten()
    |> Enum.reduce(%{}, fn x, acc -> Map.update(acc, x, 1, fn v -> v + 1 end)  end)
    |> Map.values()
    |> Enum.all?(fn x -> x == 1 end)
  end
end
