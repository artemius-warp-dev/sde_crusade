defmodule AffineCipher do
  @typedoc """
  A type for the encryption key
  """
  @type key() :: %{a: integer, b: integer}

  @m 26

  @doc """
  Encode an encrypted message using a key
  """
  @spec encode(key :: key(), message :: String.t()) :: {:ok, String.t()} | {:error, String.t()}
  def encode(%{a: a, b: _b} = keys, message) do
    coprime?(a, @m)
    |> encode(keys, message)
  end

  defp encode(coprime?, %{a: a, b: b}, message) when coprime? do
    alphabet = wrap_alphabet()

    message
    |> String.trim()
    |> String.downcase()
    |> String.replace(~r/[[:punct:]]|[[:blank:]]/i, "")
    |> String.graphemes()
    |> Enum.map(fn ch -> match_ch(ch, alphabet[ch], {a, b}, alphabet) end)
    |> Enum.chunk_every(5)
    |> Enum.map(&Enum.join/1)
    |> Enum.join(" ")
    |> then(fn res -> {:ok, res} end)
  end

  defp encode(_, _, _), do: {:error, "a and m must be coprime."}

  defp match_ch(ch, nil, _, _), do: ch

  defp match_ch(_ch, i, {a, b}, alphabet) do
    new_i = Integer.mod(a * i + b, @m)
    {encoded_ch, _} = Enum.find(alphabet, fn {_ch, j} -> j == new_i end)
    encoded_ch
  end

  defp coprime?(a, m) do
    Integer.gcd(a, m) == 1
  end

  defp wrap_alphabet() do
    ?a..?z
    |> Enum.reduce({%{}, -1}, fn ch, {acc, i} ->
      updated_i = i + 1
      {Map.put(acc, "#{[ch]}", updated_i), updated_i}
    end)
    |> then(fn {acc, _} -> acc end)
  end

  @doc """
  Decode an encrypted message using a key
  """
  @spec decode(key :: key(), message :: String.t()) :: {:ok, String.t()} | {:error, String.t()}
  def decode(%{a: a, b: _b} = keys, encrypted) do
    coprime?(a, @m)
    |> decode(keys, encrypted)
  end

  defp decode(coprime?, %{a: a, b: b}, encrypted) when coprime? do
    alphabet = wrap_alphabet()

    encrypted
    |> String.replace(~r/[[:blank:]]/, "")
    |> String.graphemes()
    |> Enum.reduce("", fn ch, acc ->
      acc <> decode_match_ch(ch, alphabet[ch], {a, b}, alphabet)
    end)
    |> then(fn str -> {:ok, str} end)
  end

  defp decode(_, _, _), do: {:error, "a and m must be coprime."}

  defp decode_match_ch(ch, nil, _, _), do: ch

  defp decode_match_ch(_ch, i, {a, b}, alphabet) do
    mmi = get_mmi(a, @m, 1)
    new_i = Integer.mod((i - b) * mmi, @m)
    {decoded_ch, _} = Enum.find(alphabet, fn {_ch, j} -> j == new_i end)
    decoded_ch
  end

  defp get_mmi(a, m, x) when rem(a * x, m) == 1, do: x
  defp get_mmi(a, m, x), do: get_mmi(a, m, x + 1)
end

# key = %{a: 5, b: 7}
# phrase = "yes"

# key = %{a: 11, b: 15}
# phrase = "mindblowingly"

# key = %{a: 17, b: 33}
# phrase = "The quick brown fox jumps over the lazy dog."

# key = %{a: 6, b: 17}
# phrase = "This is a test."

# key = %{a: 3, b: 4}
# phrase = "Testing,1 2 3, testing."

# key = %{a: 3, b: 7}
# phrase = "tytgn fjr"
## {:ok, "exercism"}

key = %{a: 13, b: 5}

phrase = "Test"

# AffineCipher.encode(key, phrase) |> IO.inspect()
AffineCipher.decode(key, phrase) |> IO.inspect()
# AffineCipher.get_mmi(15, 26, 1) |> IO.inspect()
# expected = {:ok, "xbt"}
