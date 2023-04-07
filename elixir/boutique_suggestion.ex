defmodule BoutiqueSuggestions do
  def get_combinations(tops, bottoms, options \\ []) do
    maximum_price = Keyword.get(options, :maximum_price,  100)
    for t <- tops,
        b <- bottoms,
        %{base_color: t_color} = t, 
        %{base_color: b_color} = b, t_color != b_color,
        %{price: t_price} = t,
        %{price: b_price} = b, t_price + b_price < maximum_price
      do
      {t,b}
    end  
  end
end
