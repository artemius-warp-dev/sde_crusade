defmodule Username do
  def sanitize(username) do
    sanitize(username, '')
  end

  defp sanitize([ch | tail], acc) do
    case ch do
      ch when ch >= ?a and ch <= ?z ->
        sanitize(tail, acc ++ [ch])

      ?_  ->
        sanitize(tail, acc ++ '_')

      ?ä ->
        sanitize(tail, acc ++ 'ae')

      ?ö  ->
        sanitize(tail, acc ++ 'oe')

      ?ü  ->
        sanitize(tail, acc ++ 'ue')

      ?ß  ->
        sanitize(tail, acc ++ 'ss')

      _ ->
        sanitize(tail, acc)
    end
  end

  defp sanitize([], acc), do: acc
end
