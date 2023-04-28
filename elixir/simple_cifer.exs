defmodule SimpleCipher do
  @doc """
  Given a `plaintext` and `key`, encode each character of the `plaintext` by
  shifting it by the corresponding letter in the alphabet shifted by the number
  of letters represented by the `key` character, repeating the `key` if it is
  shorter than the `plaintext`.

  For example, for the letter 'd', the alphabet is rotated to become:

  defghijklmnopqrstuvwxyzabc

  You would encode the `plaintext` by taking the current letter and mapping it
  to the letter in the same position in this rotated alphabet.

  abcdefghijklmnopqrstuvwxyz
  defghijklmnopqrstuvwxyzabc

  "a" becomes "d", "t" becomes "w", etc...

  Each letter in the `plaintext` will be encoded with the alphabet of the `key`
  character in the same position. If the `key` is shorter than the `plaintext`,
  repeat the `key`.

  Example:

  plaintext = "testing"
  key = "abc"

  The key should repeat to become the same length as the text, becoming
  "abcabca". If the key is longer than the text, only use as many letters of it
  as are necessary.
  """
  def encode(plaintext, key) do
    normalized_key = normalize(key, plaintext)
    Enum.zip(String.graphemes(plaintext), String.graphemes(normalized_key))
    |> Enum.map_join(&match/1)
  end

  @doc """
  Given a `ciphertext` and `key`, decode each character of the `ciphertext` by
  finding the corresponding letter in the alphabet shifted by the number of
  letters represented by the `key` character, repeating the `key` if it is
  shorter than the `ciphertext`.

  The same rules for key length and shifted alphabets apply as in `encode/2`,
  but you will go the opposite way, so "d" becomes "a", "w" becomes "t",
  etc..., depending on how much you shift the alphabet.
  """
  def decode(ciphertext, key) do
    oposite? = true
    normalized_key = normalize(key, ciphertext)
    Enum.zip(String.graphemes(ciphertext), String.graphemes(normalized_key))
    |> Enum.map_join(&(match(&1, oposite?)))
  end

  @doc """
  Generate a random key of a given length. It should contain lowercase letters only.
  """
  def generate_key(length) do
    alphabet = String.graphemes("abcdefghijklmnopqrstuvwxyz")
    1..length
    |> Enum.map_join(fn _ -> Enum.random(alphabet) end)
  end

  defp normalize(key, text) do
    key_length = String.length(key)
    text_length = String.length(text)
    diff = get_diff(text_length, key_length)
    
    String.duplicate(key, diff + 1)
    |> String.graphemes()
    |> Enum.take(text_length)
    |> Enum.join()
  end

  defp get_diff(a, b) when a - b < 0, do: 0
  defp get_diff(a, b), do: a - b

  defp match({ch, k}, oposite? \\ false) do
    Map.get(shift_alphabet(k, oposite?), ch)
  end
  
  defp shift_alphabet(ch, oposite?) when oposite? do
    map = shift_alphabet(ch, false)
    keys = Map.keys(map)
    values = Map.values(map)
    Enum.zip(values, keys)
    |> Map.new()
  end
  
  defp shift_alphabet(ch, _) do
    alphabet = "abcdefghijklmnopqrstuvwxyz"
    [h,t] = String.split(alphabet, ch)
    shifted_alphabet = "#{ch}#{t}#{h}"
    Enum.zip(String.graphemes(alphabet), String.graphemes(shifted_alphabet))
    |> Map.new()
  end
end
