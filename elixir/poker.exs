defmodule Poker do
  @doc """
  Given a list of poker hands, return a list containing the highest scoring hand.

  If two or more hands tie, return the list of tied hands in the order they were received.

  The basic rules and hand rankings for Poker can be found at:

  https://en.wikipedia.org/wiki/List_of_poker_hands

  For this exercise, we'll consider the game to be using no Jokers,
  so five-of-a-kind hands will not be tested. We will also consider
  the game to be using multiple decks, so it is possible for multiple
  players to have identical cards.

  Aces can be used in low (A 2 3 4 5) or high (10 J Q K A) straights, but do not count as
  a high card in the former case.

  For example, (A 2 3 4 5) will lose to (2 3 4 5 6).

  You can also assume all inputs will be valid, and do not need to perform error checking
  when parsing card values. All hands will be a list of 5 strings, containing a number
  (or letter) for the rank, followed by the suit.

  Ranks (lowest to highest): 2 3 4 5 6 7 8 9 10 J Q K A
  Suits (order doesn't matter): C D H S

  Example hand: ~w(4S 5H 4C 5D 4H) # Full house, 5s over 4s
  """
  @spec best_hand(list(list(String.t()))) :: list(list(String.t()))
  def best_hand(hands) do
    hands
    |> Enum.map(&Enum.sort/1)
    |> Enum.sort()
    |> Enum.map(&find_best_combination/1)
    |> then(fn list -> {Enum.max_by(list, fn {score, _, _} -> score end), list} end)
    |> then(fn {{maximum_score, _, _}, list} ->
      Enum.filter(list, fn {score, _, _} -> score == maximum_score end)
    end)
    |> get_maxium()
    |> Enum.map(fn {_, hand} -> hand end)
  end

  defp get_maxium(calculated_hands) do
    if same?(calculated_hands) do
      calculated_hands
      |> Enum.map(fn {score, hand, cards} ->
        {score, hand, hand -- cards}
      end)
    else
      calculated_hands
    end
    |> Enum.map(fn
      {_, [<<"2", _>>, <<"3", _>>, <<"4", _>>, <<"5", _>>, <<"A", _>>] = hand, _} ->
        {2 + 3 + 4 + 5 + 1, hand}

      {_, hand, [res_3, res_2]} when length(res_3) == 3 and length(res_2) == 2 ->
        {get_ranks_value(res_3), hand}

      {_, hand, [pair_1, pair_2]} when length(pair_1) == 2 and length(pair_2) == 2 ->
        {2 * get_ranks_value(pair_1) + get_ranks_value(pair_2) +
           get_ranks_value((hand -- pair_1) -- pair_2), hand}

      {_, hand, []} ->
        {get_highest_rank(hand) + get_ranks_value(hand), hand}

      {_, hand, cards} ->
        {Enum.reduce(hand -- cards, 0, fn card, acc ->
           acc + (match_rank(card) |> rank_to_integer())
         end)
         |> Kernel.+(get_ranks_value(cards) * 10), hand}
    end)
    |> then(fn results_with_rank ->
      {max, _} = Enum.max_by(results_with_rank, fn {rank, _} -> rank end)
      Enum.filter(results_with_rank, fn {rank, _} -> rank == max end)
    end)
  end

  defp same?(calculated_hands) do
    length(calculated_hands) !=
      length(
        Enum.uniq_by(calculated_hands, fn
          {_, _, [triplet, duplet]} when length(triplet) == 3 and length(duplet) == 2 ->
            triplet

          {_, _, cards} ->
            cards
        end)
      )
  end

  defp get_ranks_value(cards) do
    cards
    |> Enum.reduce(0, fn rank_str, acc ->
      rank = match_rank(rank_str)
      acc + rank_to_integer(rank)
    end)
  end

  defp get_highest_rank(cards) do
    cards
    |> Enum.reduce(0, &get_highest_rank/2)
  end

  defp get_highest_rank(rank_str, acc) do
    rank = match_rank(rank_str)
    integer_rank = rank_to_integer(rank)

    if acc < integer_rank do
      integer_rank
    else
      acc
    end
  end

  defp rank_to_integer(rank) do
    case rank do
      "J" -> 11
      "Q" -> 12
      "K" -> 13
      "A" -> 14
      str_number -> String.to_integer(str_number)
    end
  end

  defp find_best_combination(hand) do
    [
      {1, &high_card?/1},
      {2, &one_pair?/1},
      {3, &two_pair?/1},
      {4, &three_of_a_kind?/1},
      {5, &straight?/1},
      {6, &flash?/1},
      {7, &full_house?/1},
      {8, &four_of_a_kind?/1},
      {9, &straight_flash?/1}
    ]
    |> Enum.map(fn {score, f} -> {score, f.(hand)} end)
    |> Enum.filter(fn {_score, {res, _cards}} -> res == true end)
    |> Enum.max_by(fn {score, _} -> score end)
    |> then(fn {score, {_, cards}} -> {score, hand, cards} end)
  end

  defp straight_flash?(hand) do
    {straight?, _} = straight?(hand)
    {flash?, _} = flash?(hand)
    {straight? and flash?, hand}
  end

  defp four_of_a_kind?(hand) do
    Enum.group_by(hand, &match_rank/1)
    |> Enum.filter(fn {_key, values} -> length(values) == 4 end)
    |> then(fn
      [{_, []}] -> {false, []}
      [{_rank, list}] -> {true, list}
      _ -> {false, :not_four_of_a_kind}
    end)
  end

  defp full_house?(hand) do
    {res_3, cards_3} = three_of_a_kind?(hand)
    {res_2, cards_2} = one_pair?(hand)
    {res_3 and res_2, [cards_3, cards_2]}
  end

  defp flash?(hand) do
    suit = match_suit(hand |> hd)
    res = Enum.all?(hand, fn card -> suit == match_suit(card) end)
    {res, hand}
  end

  defp straight?(hand) do
    hand
    |> Enum.map(&(match_rank(&1) |> rank_to_integer()))
    |> Enum.sort()
    |> catch_ace_at_start()
    |> Enum.chunk_every(2, 1, :discard)
    |> Enum.all?(fn [a, b] -> b - a == 1 end)
    |> then(fn res -> {res, hand} end)
  end

  defp three_of_a_kind?(hand) do
    Enum.group_by(hand, &match_rank/1)
    |> Enum.filter(fn {_key, values} -> length(values) == 3 end)
    |> then(fn
      [{_, []}] -> {false, []}
      [{_rank, list}] -> {true, list}
      _ -> {false, :not_three_of_a_kind}
    end)
  end

  defp two_pair?(hand) do
    Enum.group_by(hand, &match_rank/1)
    |> Enum.filter(fn {_key, values} -> length(values) == 2 end)
    |> then(fn
      [{_, []}] ->
        {false, []}

      list when length(list) == 2 ->
        [{_, one_pair}, {_, second_pair}] = list
        {true, [one_pair, second_pair] |> Enum.sort(:desc)}

      _ ->
        {false, :not_two_pair}
    end)
  end

  defp high_card?(hand) do
    {true, [List.last(hand)]}
  end

  defp one_pair?(hand) do
    Enum.group_by(hand, &match_rank/1)
    |> Enum.filter(fn {_key, values} -> length(values) == 2 end)
    |> then(fn
      [{_, []}] -> {false, []}
      [{_rank, list}] -> {true, list}
      _ -> {false, :not_one_pair}
    end)
  end

  defp catch_ace_at_start([2, 3, 4, 5, 14]), do: [1, 2, 3, 4, 5]
  defp catch_ace_at_start(ranks), do: ranks

  defp match_rank(<<"10", _rest::binary>>), do: "10"
  defp match_rank(<<rank::binary-size(1), _>>), do: rank

  defp match_suit(<<"10", suit::binary>>), do: suit
  defp match_suit(<<_rank::binary-size(1), suit::binary>>), do: suit
end

# high_of_8_low_of_3 = ~w(3S 5H 6S 8D 7H)
# high_of_8_low_of_2 = ~w(2S 5D 6D 8C 7S)
# winning_hands = Poker.best_hand([high_of_8_low_of_3, high_of_8_low_of_2]) |> IO.inspect()

# high_of_king = ~w(4S 5H 6C 8D KH)
# pair_of_4 = ~w(2S 4H 6S 4D JH)
# winning_hands = Poker.best_hand([high_of_king, pair_of_4]) |> IO.inspect()

# pair_of_2 = ~w(4S 2H 6S 2D JH)
# pair_of_4 = ~w(2S 4H 6C 4D JD)
# winning_hands = Poker.best_hand([pair_of_2, pair_of_4]) |> IO.inspect()

# pair_of_8 = ~w(2S 8H 6S 8D JH)
# fives_and_fours = ~w(4S 5H 4C 8C 5C)
# winning_hands = Poker.best_hand([pair_of_8, fives_and_fours]) |> IO.inspect()

# eights_and_twos = ~w(2S 8H 2D 8D 3H)
# fives_and_fours = ~w(4S 5H 4C 8S 5D)
# winning_hands = Poker.best_hand([eights_and_twos, fives_and_fours]) |> IO.inspect()

# queens_and_twos = ~w(2S QS 2C QD JH)
# queens_and_jacks = ~w(JD QH JS 8D QC)
# winning_hands = Poker.best_hand([queens_and_twos, queens_and_jacks]) |> IO.inspect()
# queens_jacks_and_8 = ~w(JD QH JS 8D QC)
# queens_jacks_and_2 = ~w(JS QS JC 2D QD)
# winning_hands = Poker.best_hand([queens_jacks_and_8, queens_jacks_and_2]) |> IO.inspect()

# sixes_threes_and_ace = ~w(6S 6H 3S 3H AS)
# sevens_twos_and_ace = ~w(7H 7S 2H 2S AC)
# winning_hands = Poker.best_hand([sixes_threes_and_ace, sevens_twos_and_ace]) |> IO.inspect()

# fives_fours_and_two = ~w(5C 2S 5S 4H 4C)
# sixes_twos_and_seven = ~w(6S 2S 6H 7C 2C)
# winning_hands = Poker.best_hand([fives_fours_and_two, sixes_twos_and_seven]) |> IO.inspect()

# eights_and_twos = ~w(2S 8H 2H 8D JH)
# three_fours = ~w(4S 5H 4C 8S 4H)
# winning_hands = Poker.best_hand([eights_and_twos, three_fours]) |> IO.inspect()

# three_twos = ~w(2S 2H 2C 8D JH)
# three_aces = ~w(4S AH AS 8C AD)
# winning_hands = Poker.best_hand([three_twos, three_aces]) |> IO.inspect()

# three_aces_7_high = ~w(4S AH AS 7C AD)
# three_aces_8_high = ~w(4S AH AS 8C AD)
# winning_hands = Poker.best_hand([three_aces_7_high, three_aces_8_high]) |> IO.inspect()

# three_fours = ~w(4S 5H 4C 8D 4H)
# straight = ~w(3S 4D 2S 6D 5C)
# winning_hands = Poker.best_hand([three_fours, straight]) |> IO.inspect()

# three_fours = ~w(4S 5H 4C 8D 4H)
# straight_to_ace = ~w(10D JH QS KD AC)
# winning_hands = Poker.best_hand([three_fours, straight_to_ace]) |> IO.inspect()

# three_fours = ~w(4S 5H 4C 8D 4H)
# straight_to_5 = ~w(4D AH 3S 2D 5C)
# winning_hands = Poker.best_hand([three_fours, straight_to_5]) |> IO.inspect()

# pair = ~w(2C 3D 7H 5H 2S)
# not_a_straight = ~w(QS KH AC 2D 3S)
# winning_hands = Poker.best_hand([pair, not_a_straight]) |> IO.inspect()

# straight_to_8 = ~w(4S 6C 7S 8D 5H)
# straight_to_9 = ~w(5S 7H 8S 9D 6H)
# winning_hands = Poker.best_hand([straight_to_8, straight_to_9]) |> IO.inspect()

# straight_to_6 = ~w(2H 3C 4D 5D 6H)
# straight_to_5 = ~w(4S AH 3S 2D 5H)
# winning_hands = Poker.best_hand([straight_to_6, straight_to_5]) |> IO.inspect()

# straight_to_8 = ~w(4C 6H 7D 8D 5H)
# flush_to_7 = ~w(2S 4S 5S 6S 7S)
# winning_hands = Poker.best_hand([straight_to_8, flush_to_7]) |> IO.inspect()

# flush_to_9 = ~w(4H 7H 8H 9H 6H)
# flush_to_7 = ~w(2S 4S 5S 6S 7S)
# winning_hands = Poker.best_hand([flush_to_7, flush_to_9]) |> IO.inspect()

# flush_to_9_with_4_matches = ~w(3S 6S 7S 8S 9S)

# winning_hands = Poker.best_hand([flush_to_9, flush_to_9_with_4_matches]) |> IO.inspect()

# flush_to_8 = ~w(3H 6H 7H 8H 5H)
# full = ~w(4S 5H 4C 5D 4H)
# winning_hands = Poker.best_hand([flush_to_8, full]) |> IO.inspect()

# full_of_4_by_9 = ~w(4H 4S 4D 9S 9D)
# full_of_5_by_8 = ~w(5H 5S 5D 8S 8D)
# winning_hands = Poker.best_hand([full_of_4_by_9, full_of_5_by_8]) |> IO.inspect()

# full_of_5_by_9 = ~w(5H 5S 5D 9S 9D)
# full_of_5_by_8 = ~w(5H 5S 5D 8S 8D)
# winning_hands = Poker.best_hand([full_of_5_by_8, full_of_5_by_9]) |> IO.inspect()

# full = ~w(4S 5H 4D 5D 4H)
# four_3s = ~w(3S 3H 2S 3D 3C)
# winning_hands = Poker.best_hand([four_3s, full]) |> IO.inspect()

# four_2s = ~w(2S 2H 2C 8D 2D)
# four_5s = ~w(4S 5H 5S 5D 5C)
# winning_hands = Poker.best_hand([four_2s, four_5s]) |> IO.inspect()

# four_3s_and_2 = ~w(3S 3H 2S 3D 3C)
# four_3s_and_4 = ~w(3S 3H 4S 3D 3C)
# winning_hands = Poker.best_hand([four_3s_and_2, four_3s_and_4]) |> IO.inspect()

# four_5s = ~w(4S 5H 5S 5D 5C)
# straight_flush_to_10 = ~w(7S 8S 9S 6S 10S)
# winning_hands = Poker.best_hand([four_5s, straight_flush_to_10]) |> IO.inspect()

# four_aces = ~w(KC AH AS AD AC)
# straight_flush_to_ace = ~w(10C JC QC KC AC)
# winning_hands = Poker.best_hand([four_aces, straight_flush_to_ace]) |> IO.inspect()

# four_aces = ~w(KS AH AS AD AC)
# straight_flush_to_5 = ~w(4H AH 3H 2H 5H)
# winning_hands = Poker.best_hand([four_aces, straight_flush_to_5]) |> IO.inspect()

# nothing = ~w(2C AC QC 10C KC)
# not_a_straight_flush = ~w(QH KH AH 2H 3H)
# winning_hands = Poker.best_hand([nothing, not_a_straight_flush]) |> IO.inspect()

# straight_flush_to_8 = ~w(4H 6H 7H 8H 5H)
# straight_flush_to_9 = ~w(5S 7S 8S 9S 6S)
# winning_hands = Poker.best_hand([straight_flush_to_8, straight_flush_to_9]) |> IO.inspect()

# straight_flush_to_6 = ~w(2H 3H 4H 5H 6H)
# straight_flush_to_5 = ~w(4D AD 3D 2D 5D)
# winning_hands = Poker.best_hand([straight_flush_to_6, straight_flush_to_5]) |> IO.inspect()

high_of_8 = ~w(4D 5S 6S 8D 3C)
high_of_10 = ~w(2S 4C 7S 9H 10H)
high_of_jack = ~w(3S 4S 5D 6H JH)
another_high_of_jack = ~w(3H 4H 5C 6C JD)
hands = [high_of_8, high_of_10, high_of_jack, another_high_of_jack]
winning_hands = Poker.best_hand(hands) |> IO.inspect()

assert_poker = fn winning_hands, expected_winning_hands ->
  IO.inspect({winning_hands, expected_winning_hands})

  winning_hands =
    if is_list(winning_hands) do
      winning_hands
      |> Enum.map(&Enum.sort/1)
      |> Enum.sort()
    else
      winning_hands
    end

  expected_winning_hands =
    if is_list(expected_winning_hands) do
      expected_winning_hands
      |> Enum.map(&Enum.sort/1)
      |> Enum.sort()
    else
      expected_winning_hands
    end

  (winning_hands == expected_winning_hands) |> IO.inspect()
end

# high_of_jack = ~w(4S 5S 7H 8D JC)

# winning_hands = Poker.best_hand([high_of_jack])
# assert_poker.(winning_hands, [high_of_jack])
