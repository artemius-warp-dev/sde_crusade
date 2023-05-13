defmodule Diamond do
  @doc """
  Given a letter, it prints a diamond starting with 'A',
  with the supplied letter at the widest point.
  """
  @spec build_shape(char) :: String.t()
  def build_shape(letter) do
    half = build_shape(?A, letter, [], letter - ?A)
    [_ | tail] = half

    (Enum.reverse(half) ++ tail)
    |> Enum.join()
  end

  defp build_shape(?A, _acc, _, 0), do: ["A\n"]

  defp build_shape(?A, end_letter, acc, spaces) do
    updated_acc = [
      "#{String.duplicate("\s", spaces)}#{[?A]}#{String.duplicate("\s", spaces)}\n"
      | acc
    ]

    build_shape(?A + 1, end_letter, updated_acc, spaces - 1)
  end

  defp build_shape(end_letter, end_letter, acc, 0) do
    amount = (end_letter - ?A) * 2 - 1
    ["#{[end_letter]}#{String.duplicate("\s", amount)}#{[end_letter]}\n" | acc]
  end

  defp build_shape(letter, end_letter, acc, spaces) do
    amount = (end_letter - ?A) * 2 - 1

    updated_acc = [
      "#{String.duplicate("\s", spaces)}#{[letter]}#{String.duplicate("\s", amount - 2 * spaces)}#{[letter]}#{String.duplicate("\s", spaces)}\n"
      | acc
    ]

    build_shape(letter + 1, end_letter, updated_acc, spaces - 1)
  end
end

Diamond.build_shape(?D) |> IO.inspect()
