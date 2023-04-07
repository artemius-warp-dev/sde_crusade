defmodule Acronym do
  @doc """
  Generate an acronym from a string.
  "This is a string" => "TIAS"
  """
  @spec abbreviate(String.t()) :: String.t()
  def abbreviate(string) do
    string
    |> String.replace("_", "")
    |> String.trim()
    |> String.split([" ", "-"])
    |> Enum.filter(fn x -> x != "" end)
    |> Enum.map(&String.first(&1) |> String.upcase() )
    |> Enum.join()
  end
end
