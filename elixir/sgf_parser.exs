defmodule SgfParsing do
  defmodule Sgf do
    defstruct properties: %{}, children: []
  end

  @type sgf :: %Sgf{properties: map, children: [sgf]}

  @doc """
  Parse a string into a Smart Game Format tree
  """
  @spec parse(encoded :: String.t()) :: {:ok, sgf} | {:error, String.t()}
  def parse(encoded) do
    cond do
      encoded == "" -> {:error, "tree missing"}
      encoded == "()" -> {:error, "tree with no nodes"}
      encoded == ";" -> {:error, "tree missing"}
      encoded == "(;)" -> {:ok, %SgfParsing.Sgf{}}
      true -> do_parse(encoded)
    end
  end

  defp ensure_delimiters([]), do: throw("properties without delimiter")

  defp ensure_delimiters(nodes), do: nodes

  defp ensure_upcase_properties(encoded) do
    Regex.scan(~r/(\w)\[/, encoded, capture: :all_but_first)
    |> List.flatten()
    |> Enum.all?(fn x -> x == String.upcase(x) end)
    |> then(fn
      true -> encoded
      false -> throw("property must be in uppercase")
    end)
  end

  defp do_parse(encoded) do
    processed_encoded =
      encoded
      |> String.replace(~r/\\\t/, " ")
      |> String.replace(~r/\\\n/, "")
      |> String.replace(~r/\\t/, "t")
      |> String.replace(~r/\\n/, "n")
      |> String.replace(~r/\t/, " ")
      |> String.replace("\\\\", "**")
      |> String.replace(~r/\[(\w+)\[/, "[\\1\\[")
      |> String.replace("\\[", "<")
      |> String.replace("\\]", ">")
      |> String.replace("][", ",")

    try do
      processed_encoded
      |> ensure_upcase_properties()
      |> then(&Regex.scan(~r/(;?)([A-Z]+)\[([^\[\]]+)\]/, &1))
      |> ensure_delimiters
      |> Enum.map(fn [_, tree, key, value] -> [tree, Map.new([{key, normalize(value)}])] end)
      |> Enum.reduce(%{children: []}, &merge/2)
      |> then(&{:ok, struct(SgfParsing.Sgf, &1)})
    catch
      error ->
        {:error, error}
    end
  end

  defp normalize(value) do
    value
    |> String.replace("**", "\\")
    |> String.replace("<", "\[")
    |> String.replace(">", "\]")
    |> String.replace("\\n", "\n")
    |> String.replace("\\t", " ")
    |> String.split(",")
  end

  defp merge([tree, map], acc) do
    cond do
      tree == ";" && !acc[:properties] -> Map.merge(acc, %{properties: map})
      tree == "" -> %{acc | properties: Map.merge(acc.properties, map)}
      true -> %{acc | children: acc.children ++ [struct(SgfParsing.Sgf, %{properties: map})]}
    end
  end
end

# encoded = "(;A[B](;B[C])(;C[D]))"
# encoded = "(;A[hello\t\tworld])"
# encoded = "(;A[x[y\\]z][foo]B[bar];C[baz])"
encoded = "(;A)"
# encoded = "(;a[b])"
# {:ok,
#  %Sgf{
#    children: [
#      %Sgf{properties: %{"B" => ["C"]}},
#      %Sgf{properties: %{"C" => ["D"]}}
#    ],
#    properties: %{"A" => ["B"]}
#  }}

# encoded = "(;A[b][c][d])"
# expected = {:ok, %Sgf{properties: %{"A" => ["b", "c", "d"]}}}

SgfParsing.parse(encoded) |> IO.inspect()
