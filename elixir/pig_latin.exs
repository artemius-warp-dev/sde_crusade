defmodule PigLatin do
  @doc """
  Given a `phrase`, translate it a word at a time to Pig Latin.
  """

  @vowels ['e', 'y', 'u', 'i', 'o', 'a']
  @consonants [
    'q',
    'w',
    'r',
    't',
    'p',
    's',
    'd',
    'f',
    'g',
    'h',
    'j',
    'k',
    'l',
    'z',
    'x',
    'c',
    'v',
    'b',
    'n',
    'm'
  ]

  @spec translate(phrase :: String.t()) :: String.t()
  def translate(phrase) do
    phrase
    |> String.split(" ")
    |> Enum.reduce([], &do_rules/2)
    |> Enum.reverse()
    |> Enum.join(" ")
  end

  defp do_rules(word, acc) do
    updated_word =
      word
      |> String.trim()
      |> do_rule_1()

    [updated_word | acc]
  end

  defp do_rule_1(<<f::size(8), s::size(8), _rest::bitstring>> = str)
       when [s] in @vowels and [f] == 'y',
       do: do_rule_2_3_4(str)

  defp do_rule_1(<<f::size(8), s::size(8), _rest::bitstring>> = str)
       when [f] in @vowels or [f, s] in ['xb', 'xr', 'yt'] do
    str <> "ay"
  end

  defp do_rule_1(word), do: do_rule_2_3_4(word)

  defp find_vowel(<<>>, acc), do: {<<>>, acc}
  defp find_vowel(<<ch::8, _rest::bitstring>> = str, acc) when [ch] in @vowels, do: {str, acc}
  defp find_vowel(<<ch::8, rest::bitstring>>, acc), do: find_vowel(rest, [ch | acc])

  defp do_rule_2_3_4(<<f::size(8), _rest::bitstring>> = str)
       when [f] in @consonants or [f] == 'y' do
    {res, acc} = find_vowel(str, [])
    acc_str = Enum.reverse(acc) |> to_string()

    case {String.last(acc_str), String.first(res)} do
      {"q", "u"} -> "#{String.slice(res, 1..String.length(res))}#{String.trim(acc_str, "q")}quay"
      {nil, "y"} -> "#{String.slice(res, 1..String.length(res))}yay"
      {_, "y"} -> "y#{String.slice(res, 1..String.length(res))}#{acc_str}ay"
      _ -> "#{res}#{acc_str}ay"
    end
  end

  defp do_rule_2_3_4(word), do: word
end
