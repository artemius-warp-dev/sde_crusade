defmodule Allergies do
  @doc """
  List the allergies for which the corresponding flag bit is true.
  """

  @items %{
    1 => "eggs",
    2 => "peanuts",
    4 => "shellfish",
    8 => "strawberries",
    16 => "tomatoes",
    32 => "chocolate",
    64 => "pollen",
    128 => "cats"
  }

  @spec list(non_neg_integer) :: [String.t()]
  def list(flags) do
    get_score(flags, [], 128)
    |> Enum.map(&@items[&1])
  end

  defp get_score(0, acc, _), do: acc

  defp get_score(flag, acc, weight) when flag >= 256, do: get_score(flag - 256, acc, weight)

  defp get_score(flag, acc, weight) when flag - weight >= 0 and weight <= 128 do
    get_score(flag - weight, [weight | acc], step(div(weight, 2)))
  end

  defp get_score(flag, acc, weight), do: get_score(flag, acc, step(div(weight, 2)))

  defp step(0), do: 1
  defp step(weight), do: weight

  @doc """
  Returns whether the corresponding flag bit in 'flags' is set for the item.
  """
  @spec allergic_to?(non_neg_integer, String.t()) :: boolean
  def allergic_to?(flags, item) do
    item in list(flags)
  end

  #  @allergies %{
  #     "eggs" => 1,
  #     "peanuts" => 2,
  #     "shellfish" => 4,
  #     "strawberries" => 8,
  #     "tomatoes" => 16,
  #     "chocolate" => 32,
  #     "pollen" => 64,
  #     "cats" => 128
  #   }

  #   @doc """

  #   List the allergies for which the corresponding flag bit is true.

  #   """

  #   @spec list(non_neg_integer) :: [String.t()]

  #   def list(flags) do
  #     @allergies
  #     |> Map.keys()
  #     |> Enum.filter(&allergic_to?(flags, &1))
  #   end

  #   @doc """

  #   Returns whether the corresponding flag bit in 'flags' is set for the item.

  #   """

  #   @spec allergic_to?(non_neg_integer, String.t()) :: boolean

  #   def allergic_to?(flags, item) do
  #     (@allergies[item] &&& flags) > 0
  #   end
end

# 248
Allergies.list(248) |> IO.inspect()
