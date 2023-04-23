defmodule WordCount do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  # @spec count(String.t()) :: map
  # def count(sentence) do
  #   sentence
  #   |> String.downcase()
  #   |> String.replace([":", "!", "?", "&", "@", "$", "%", "^", ".", "'"], "")
  #   |> String.trim()
  #   |> String.split([" ", ",", "\n", "_"])
  #   |> Enum.filter(&(&1 != ""))
  #   |> Enum.reduce({nil, %{}}, fn word, {_, acc} ->
  #     Map.get_and_update(acc, to_contractions(word), fn
  #       nil -> {nil, 1}
  #       v -> {v, v + 1}
  #     end)
  #   end)
  #   |> then(fn {_, map} -> map end)
  # end

  # defp to_contractions(word) do
  #   case word do
  #     "dont" -> "don't"
  #     "youre" -> "you're"
  #     "cant" -> "can't"
  #     _ -> word
  #   end
  # end

  @spec count(String.t()) :: map
  def count(phrase) do
    phrase
    |> String.downcase()
    |> to_words()
    |> Enum.frequencies()
  end

  defp to_words(phrase) do
    ~r/[\s_,!&@$%^:.]+/
    |> Regex.split(phrase)
    |> Enum.reject(&(&1 == ""))
    |> Enum.map(&normalize_word/1)
  end

  defp normalize_word(word) do
    word |> String.trim_leading("'") |> String.trim_trailing("'")
  end
end
