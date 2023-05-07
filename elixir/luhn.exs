defmodule Luhn do
  @doc """
  Checks if the given number is valid via the luhn formula
  """
  @spec valid?(String.t()) :: boolean
  def valid?(number) when is_binary(number) do
      with trimmed_str = String.trim(number),
           true <- not String.match?(trimmed_str, ~r/[^[:digit:][:blank:]]/) |> IO.inspect(),
           true <- String.length(trimmed_str) > 1,
           strs = Regex.scan(~r/[[:digit:]]/, trimmed_str) |> List.flatten(),
           numbers = Enum.map(strs, &String.to_integer/1) do
      valid?(numbers) 
      end
  end

  def valid?(numbers) when is_list(numbers) do
    numbers
    |> Enum.reverse()
    |> Enum.reduce({1, []}, &transform/2)
    |> then(fn {_, acc} -> acc end)
    |> Enum.sum()
    |> then(fn x -> rem(x, 10) == 0 end)
  end

  defp double_x(x) when x * 2 > 9, do: x*2 - 9
  defp double_x(x), do:  x * 2

  defp transform(x, {i, acc}) when rem(i, 2) == 0 do
    {i+1, [double_x(x) | acc]}
  end
  defp transform(x, {i, acc}) do
    {i+1, [x | acc]}
  end

end

Luhn.valid?("059a") |> IO.inspect()
