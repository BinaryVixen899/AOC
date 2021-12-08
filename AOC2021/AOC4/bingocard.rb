class BingoCard
  attr_reader :stringofnumbers, :dictionary, :solvabledictionary
  
  #mark it 
#check all rows and all columns for marks to see if they've won
  def parse(card)
    #I know i can do this in parse, I need to split these up into 5 a piece and then create rows out of them 
  end

  def initialize(bingocardnumbers)
    # if bingocardnumbers != Array
    #   puts "Please provide an array"
    # end
    @stringofnumbers = bingocardnumbers 
    newstringofnumbers = []
    @dictionary = {}
    stringofnumbers.each_slice(5) {|x| newstringofnumbers.push(x) }
    stringofnumbers.each_index{|x| dictionary[stringofnumbers[x]] = x }
    @solvabledictionary = dictionary
  end

  def mark_it(inputnumbers)
    # if inputnumbers != Array
    #   puts "Please provide an array"
    # end
    inputnumbers.each_with_index do  |x, i| 
      if dictionary.has_key?(x) 
        puts x, dictionary[x]
        solvabledictionary[x] = true
        
      else
        solvabledictionary[x] = false
      end
    # we also need to solve columns here at some point 
  end
  
  def check_rows
    truecount = 0 
    

   
      stringofnumbers.each do |x| 
        puts x
        if solvabledictionary[x] == true && truecount < 5
          truecount += 1 
        elsif truecount == 5
          puts "I win!"
          break
          # move this to first 
        elsif truecount != 5
          truecount = 0 
        end
      
      end

    puts truecount

    if truecount == 5
      puts "This card is a winner!"
    else
      puts "Not such a winner!"
    end

  end

  
end



  
end

# mycard = BingoCard.new([22, 13, 17, 11, 0, 8,  2, 23,  4, 24, 21,  9, 14, 16,  7, 6, 10,  3, 18,  5, 1, 12, 20, 15, 19])
mycard = BingoCard.new([14, 21, 17, 24,  4, 10, 16, 15,  9, 19, 18, 8, 23, 26, 20, 22, 11, 13, 6,  5, 2,  0, 12,  3,  7])
puts mycard.solvabledictionary
mycard.mark_it([7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1])
puts mycard.solvabledictionary
mycard.check_rows()
# Bingo
# inputnumbers = 
