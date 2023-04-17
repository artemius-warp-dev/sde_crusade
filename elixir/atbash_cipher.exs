defmodule Atbash do
  @doc """
  Encode a given plaintext to the corresponding ciphertext
  ## Examples

  iex> Atbash.encode("completely insecure")
  "xlnko vgvob rmhvx fiv"
  """

  @cache Enum.zip(
    String.codepoints("abcdefghijklmnopqrstuvwxyz"), 
    String.codepoints("zyxwvutsrqponmlkjihgfedcba")
    ) |> Map.new()

  @spec encode(String.t()) :: String.t()
  def encode(plaintext) do
    plaintext
    |> String.downcase()
    |> String.codepoints()
    |> Enum.map(&match/1)
    |> Enum.filter(&(&1 != ""))
    |> Enum.chunk_every(5)
    |> Enum.map(&Enum.join/1)
    |> Enum.join(" ")
  end

  defp match(ch) do
    
    cond do
      String.contains?("0123456789", ch) -> ch
      String.match?(ch, ~r/[[:punct:]]/) -> ""
      ch == " " -> ""
      true -> @cache[ch]
    end
  end

  @spec decode(String.t()) :: String.t()
  def decode(cipher) do
    cipher
    |> String.codepoints()
    |> Enum.filter(&(&1 != " "))
    |> Enum.map(&match/1)
    |> Enum.join()
  end
end
