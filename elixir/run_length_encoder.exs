defmodule RunLengthEncoder do
  @doc """
  Generates a string where consecutive elements are represented as a data value and count.
  "AABBBCCCC" => "2A3B4C"
  For this example, assume all input are strings, that are all uppercase letters.
  It should also be able to reconstruct the data into its original form.
  "2A3B4C" => "AABBBCCCC"
  """
  @spec encode(String.t()) :: String.t()
  def encode(string) do
    string
    |> String.graphemes()
    |> Enum.reduce([], &to_cache/2)
    |> Enum.reverse()
    |> Enum.map(&from_cache/1)
    |> Enum.join()
  end

  @spec decode(String.t()) :: String.t()
  def decode(string) do
    string
    |> String.graphemes()
    |> Enum.reduce({"", "0"}, &decode/2)
    |> then(fn {str, _} -> str end)
  end

  defp decode(ch, {str, count_str}) do
    parsed_count = String.to_integer(count_str)

    case Integer.parse(ch) do
      {c, ""} ->
        {str, "#{count_str}#{c}"}

      _ when parsed_count > 1 ->
        {"#{str}#{1..parsed_count |> Enum.map_join(fn _ -> ch end)}", "0"}

      _ ->
        {"#{str}#{ch}", "0"}
    end
  end

  @spec encode(String.t()) :: String.t()
  def encode(string) do
    string
    |> String.codepoints()
    |> Enum.chunk_by(& &1)
    |> Enum.map(&compress/1)
    |> Enum.join()
  end

  defp compress([ch]), do: ch
  defp compress(list), do: "#{length(list)}#{List.first(list)}"

  @spec decode(String.t()) :: String.t()
  def decode(string) do
    Regex.scan(~r/(\d*)(.)/, string)
    |> Enum.map(&decompress/1)
    |> Enum.join()
  end

  defp decompress([_, "", ch]), do: ch
  defp decompress([_, count_str, ch]), do: String.duplicate(ch, String.to_integer(count_str))

  defp from_cache({key, value}) when value > 1, do: "#{value}#{key}"
  defp from_cache({key, _value}), do: "#{key}"

  defp to_cache(ch, [{prev_ch, count} | tail]) when ch == prev_ch do
    [{ch, count + 1} | tail]
  end

  defp to_cache(ch, acc), do: [{ch, 1} | acc]
end
