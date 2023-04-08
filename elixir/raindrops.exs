defmodule Raindrops do
  @doc """
  Returns a string based on raindrop factors.

  - If the number contains 3 as a prime factor, output 'Pling'.
  - If the number contains 5 as a prime factor, output 'Plang'.
  - If the number contains 7 as a prime factor, output 'Plong'.
  - If the number does not contain 3, 5, or 7 as a prime factor,
    just pass the number's digits straight through.
  """
  @spec convert(pos_integer) :: String.t()
  def convert(number) do
    ""
    |> add_pling(rem(number, 3) == 0)
    |> add_plang(rem(number, 5) == 0)
    |> add_plong(rem(number, 7) == 0)
    |> wrap_result(number)
  end

  defp add_pling(_result, factor3?) when factor3?, do: "Pling"
  defp add_pling(result, _), do: result
  defp add_plang(result, factor5?) when factor5?, do: "#{result}Plang"
  defp add_plang(result, _), do: result
  defp add_plong(result, factor7?) when factor7?, do: "#{result}Plong"
  defp add_plong(result, _), do: result

  defp wrap_result("", number), do: to_string(number)
  defp wrap_result(result, _), do: result
end
