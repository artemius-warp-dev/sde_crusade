defmodule Bob do
  @spec hey(String.t()) :: String.t()
  def hey(str) do
    input = String.trim(str)

    cond do
      String.trim(input) == "" ->
        "Fine. Be that way!"

      yell?(input) and String.last(input) == "?" ->
        "Calm down, I know what I'm doing!"

      yell?(input) ->
        "Whoa, chill out!"

      String.last(input) == "?" ->
        "Sure."

      true ->
        "Whatever."
    end
  end

  defp yell?(input) do
    input == String.upcase(input) && Regex.match?(~r/[[:alpha:]]/, input)
  end
end

Bob.hey("WATCH OUT!") |> IO.inspect()
