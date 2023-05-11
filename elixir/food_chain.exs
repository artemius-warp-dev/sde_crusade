defmodule FoodChain do
  @doc """
  Generate consecutive verses of the song 'I Know an Old Lady Who Swallowed a Fly'.
  """

  @lady_swallowed "I know an old lady who swallowed"

  @spec recite(start :: integer, stop :: integer) :: String.t()
  def recite(start, stop) do
    start..stop
    |> Enum.map(&recite/1)
    |> Enum.join("\n")
  end

  defp combine(to_start, amount) do
    (to_start ++ to_end(amount))
    |> Enum.join("\n")
  end

  defp to_end(amount) do
    [
      "I don't know why she swallowed the fly. Perhaps she'll die.\n",
      "She swallowed the spider to catch the fly.",
      "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
      "She swallowed the cat to catch the bird.",
      "She swallowed the dog to catch the cat.",
      "She swallowed the goat to catch the dog.",
      "She swallowed the cow to catch the goat."
    ]
    |> Enum.take(amount)
    |> Enum.reverse()
  end

  defp recite(1) do
    ["#{@lady_swallowed} a fly."]
    |> combine(1)
  end

  defp recite(2) do
    ["#{@lady_swallowed} a spider.", "It wriggled and jiggled and tickled inside her."]
    |> combine(2)
  end

  defp recite(3) do
    [
      "#{@lady_swallowed} a bird.",
      "How absurd to swallow a bird!"
    ]
    |> combine(3)
  end

  defp recite(4) do
    [
      "#{@lady_swallowed} a cat.",
      "Imagine that, to swallow a cat!"
    ]
    |> combine(4)
  end

  defp recite(5) do
    [
      "#{@lady_swallowed} a dog.",
      "What a hog, to swallow a dog!"
    ]
    |> combine(5)
  end

  defp recite(6) do
    [
      "#{@lady_swallowed} a goat.",
      "Just opened her throat and swallowed a goat!"
    ]
    |> combine(6)
  end

  defp recite(7) do
    [
      "#{@lady_swallowed} a cow.",
      "I don't know how she swallowed a cow!"
    ]
    |> combine(7)
  end

  defp recite(8) do
    [
      "#{@lady_swallowed} a horse.",
      "She's dead, of course!\n"
    ]
    |> Enum.join("\n")
  end
end

# defmodule FoodChain do
#   @animals [
#     fly: "I don't know why she swallowed the fly.",
#     spider: "It wriggled and jiggled and tickled inside her.",
#     bird: "How absurd to swallow a bird!",
#     cat: "Imagine that, to swallow a cat!",
#     dog: "What a hog, to swallow a dog!",
#     goat: "Just opened her throat and swallowed a goat!",
#     cow: "I don't know how she swallowed a cow!",
#     horse: "She's dead, of course!"
#   ]

#   @doc """
#   Generate consecutive verses of the song 'I Know an Old Lady Who Swallowed a Fly'.
#   """

#   @spec recite(start :: integer, stop :: integer) :: String.t()

#   def recite(start, stop), do: Enum.map_join(start..stop, "\n", &verse/1)

#   def verse(1),
#     do: """
#     I know an old lady who swallowed a fly.
#     I don't know why she swallowed the fly. Perhaps she'll die.
#     """

#   def verse(8),
#     do: """
#     I know an old lady who swallowed a horse.
#     She's dead, of course!
#     """

#   def verse(number) do
#     [{animal, line} | _] =
#       animals =
#       Enum.take(@animals, number)
#       |> Enum.reverse()

#     cumulative_lines =
#       Enum.zip(animals, tl(animals))
#       |> Enum.map_join("\n", fn
#         {{:bird, _}, _} ->
#           "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her."

#         {{a, _}, {b, _}} ->
#           "She swallowed the #{a} to catch the #{b}."
#       end)

#     """
#     I know an old lady who swallowed a #{animal}.
#     #{line}
#     #{cumulative_lines}
#     I don't know why she swallowed the fly. Perhaps she'll die.
#     """
#   end
# end
