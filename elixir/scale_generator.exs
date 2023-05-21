defmodule ScaleGenerator do
  @doc """
  Find the note for a given interval (`step`) in a `scale` after the `tonic`.

  "m": one semitone
  "M": two semitones (full tone)
  "A": augmented second (three semitones)

  Given the `tonic` "D" in the `scale` (C C# D D# E F F# G G# A A# B C), you
  should return the following notes for the given `step`:

  "m": D#
  "M": E
  "A": F
  """
  @spec step(scale :: list(String.t()), tonic :: String.t(), step :: String.t()) :: String.t()
  def step(scale, tonic, step) do
    tonic_index = Enum.find_index(scale, &(&1 == String.capitalize(tonic)))

    case step do
      "m" -> 1
      "M" -> 2
      "A" -> 3
    end
    |> then(fn i -> Enum.at(scale, tonic_index + i) end)
  end

  @doc """
  The chromatic scale is a musical scale with thirteen pitches, each a semitone
  (half-tone) above or below another.

  Notes with a sharp (#) are a semitone higher than the note below them, where
  the next letter note is a full tone except in the case of B and E, which have
  no sharps.

  Generate these notes, starting with the given `tonic` and wrapping back
  around to the note before it, ending with the tonic an octave higher than the
  original. If the `tonic` is lowercase, capitalize it.

  "C" should generate: ~w(C C# D D# E F F# G G# A A# B C)
  """
  @spec chromatic_scale(tonic :: String.t()) :: list(String.t())
  def chromatic_scale(tonic \\ "C") do
    tonic
    |> String.capitalize()
    |> chromatic_scale(0, [])
  end

  defp chromatic_scale(_tonic, 13, acc), do: Enum.reverse(acc)

  defp chromatic_scale(<<ch::size(8), "#">> = str, count, acc) do
    chromatic_scale(ch + 1, count + 1, [str | acc])
  end

  defp chromatic_scale(<<ch::size(8)>>, count, acc) do
    chromatic_scale(ch, count, acc)
  end

  defp chromatic_scale(ch, count, acc) when ch in [?E, ?B],
    do: chromatic_scale(ch + 1, count + 1, [to_string([ch]) | acc])

  defp chromatic_scale(ch, count, acc) do
    {str, next_ch} = choose_ch(ch, acc)
    chromatic_scale(next_ch, count + 1, [str | acc])
  end

  defp choose_ch(ch, []), do: {[ch] |> to_string(), ch}

  defp choose_ch(ch, acc) when ch > ?G do
    choose_ch(?A, acc)
  end

  defp choose_ch(ch, [h | _]) do
    cond do
      ch in [?E, ?B] -> {[ch] |> to_string(), ch + 1}
      String.ends_with?(h, "#") or h in ["E", "B"] -> {[ch] |> to_string(), ch}
      true -> {to_string([ch]) <> "#", ch + 1}
    end
  end

  @doc """
  Sharp notes can also be considered the flat (b) note of the tone above them,
  so the notes can also be represented as:

  A Bb B C Db D Eb E F Gb G Ab

  Generate these notes, starting with the given `tonic` and wrapping back
  around to the note before it, ending with the tonic an octave higher than the
  original. If the `tonic` is lowercase, capitalize it.

  "C" should generate: ~w(C Db D Eb E F Gb G Ab A Bb B C)
  """
  @spec flat_chromatic_scale(tonic :: String.t()) :: list(String.t())
  def flat_chromatic_scale(tonic \\ "C") do
    tonic
    |> String.capitalize()
    |> flat_chromatic_scale(0, [])
  end

  defp flat_chromatic_scale(_tonic, 13, acc), do: Enum.reverse(acc)

  defp flat_chromatic_scale(ch, count, acc) when ch in [?F, ?C],
    do: flat_chromatic_scale(ch + 1, count + 1, [to_string([ch]) | acc])

  defp flat_chromatic_scale(ch, count, acc) do
    {str, next_ch} = flat_choose_ch(ch, acc)
    flat_chromatic_scale(next_ch, count + 1, [str | acc])
  end

  defp flat_choose_ch(<<ch::size(8), "b">> = str, _acc), do: {str, ch}
  defp flat_choose_ch(<<ch::size(8)>>, []), do: {[ch] |> to_string(), ch + 1}

  defp flat_choose_ch(ch, acc) when ch > ?G do
    flat_choose_ch(?A, acc)
  end

  defp flat_choose_ch(ch, [h | _]) do
    cond do
      String.ends_with?(h, "b") -> {[ch] |> to_string(), ch + 1}
      true -> {to_string([ch]) <> "b", ch}
    end
  end

  @doc """
  Certain scales will require the use of the flat version, depending on the
  `tonic` (key) that begins them, which is C in the above examples.

  For any of the following tonics, use the flat chromatic scale:

  F Bb Eb Ab Db Gb d g c f bb eb

  For all others, use the regular chromatic scale.
  """
  @spec find_chromatic_scale(tonic :: String.t()) :: list(String.t())
  def find_chromatic_scale(tonic) do
    cond do
      tonic in ~w(F Bb Eb Ab Db Gb d g c f bb eb) -> flat_chromatic_scale(tonic)
      true -> chromatic_scale(tonic)
    end
  end

  @doc """
  The `pattern` string will let you know how many steps to make for the next
  note in the scale.

  For example, a C Major scale will receive the pattern "MMmMMMm", which
  indicates you will start with C, make a full step over C# to D, another over
  D# to E, then a semitone, stepping from E to F (again, E has no sharp). You
  can follow the rest of the pattern to get:

  C D E F G A B C
  """
  @spec scale(tonic :: String.t(), pattern :: String.t()) :: list(String.t())
  def scale(tonic, pattern) do
    scale = find_chromatic_scale(tonic)

    pattern
    |> String.graphemes()
    |> Enum.reduce([String.capitalize(tonic)], fn step, [h | _t] = acc ->
      [step(scale, h, step) | acc]
    end)
    |> Enum.reverse()
  end
end

ScaleGenerator.step(
  ["C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C"],
  "D",
  "M"
)
|> IO.inspect() == ~w(C D E F G A B C)
