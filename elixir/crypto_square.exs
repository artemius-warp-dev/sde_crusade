defmodule CryptoSquare do
  @doc """
  Encode string square methods
  ## Examples

    iex> CryptoSquare.encode("abcd")
    "ac bd"
  """
  @spec encode(String.t()) :: String.t()
  def encode(str) do
    str
    |> String.replace(~r/[^[:alnum:]]/, "")
    |> String.downcase()
    |> do_encode()
  end

  defp do_encode(""), do: ""

  defp do_encode(normalized_str) do
    [r, c] = find_c_and_r(String.length(normalized_str))

    chunked_list =
      1..r
      |> Enum.reduce({[], normalized_str}, fn x, {acc, remaind_str} ->
        {str, updated_remaind_str} = String.split_at(remaind_str, c)
        normalized_chunk_str = String.pad_trailing(str, c)
        {[normalized_chunk_str | acc], updated_remaind_str}
      end)
      |> then(fn {acc, _} -> Enum.reverse(acc) end)

    1..c
    |> Enum.map(&generate_str(&1, chunked_list))
    |> Enum.join(" ")
  end

  defp generate_str(i, list) do
    list
    |> Enum.reduce("", fn str, acc -> acc <> String.at(str, i - 1) end)
  end

  defp find_c_and_r(length) do
    for x <- -1..length, y <- 1..length do
      [x, y]
    end
    |> Enum.find(&find_c_and_r(&1, length))
  end

  defp find_c_and_r([r, c], length) do
    r * c >= length and c >= r and c - r <= 1
  end
end

sentence = "A"
CryptoSquare.encode(sentence) |> IO.inspect(charlists: :as_lists)
