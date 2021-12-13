require 'pry'
class BingoCard
  attr_reader :bingocardnumbers, :solvabledictionary, :score, :winningnumber, :lastnumber, :calledarray, :sum, :name, :turnstowin
  
  #mark it 
#check all rows and all columns for marks to see if they've won
  def parse(card)
    #I know i can do this in parse, I need to split these up into 5 a piece and then create rows out of them 
  end

  def initialize(name, bingocardnumbers, inputnumbers)
    # if bingocardnumbers != Array
    #   puts "Please provide an array"
    # end
    @bingocardnumbers = bingocardnumbers 
    @inputnumbers = inputnumbers
    @solvabledictionary = {}
    @calledarray = []
    # A string of numbers that are our bingo card numbers
    #for some reason a new string of numbers
    bingocardnumbers.each_index{|x| @solvabledictionary[bingocardnumbers[x]] = x }
  end

  def mark_it(inputnumbers)
    won = nil
    # if inputnumbers != Array
    #   puts "Please provide an array"
    # end
    inputnumbers.each_with_index do |x, i| 
        #iterates over bingo card numbers 
        
        #adds them to an array 
        # Serena is right, I should sleep. But if I want an even more cursed an inificient answer I could just, you know. add every single x to an array and subtract by one

        if won == true
          #checks if won is true
          break
        else
          @calledarray.push(x)
          @solvabledictionary[x] = true
        end
        
        won = check_if_bingo(i) 
        # above is horizontal bingo 

          
        # binding.pry
        
        #if won isn't true takes the solveable dictionary for input numbers and sets to true 
        
        
      end
    # we also need to solve columns here at some point \
    #So what we do here is we iterate over every input number and then check if we have five in a row 
  end

  def check_if_bingo(index)
    yescount = 0
    dictionarylist = []
    # maybe a while i is less than 25
    @solvabledictionary.each.with_index(1) do |(key, value), i|
      # 4 is the number that of numbers that we want here 
      # we want to iterate 4 at a time and see if the value of all 4 is true
      #if it is then we call it
      #if it isn't then we go back to the main program
    

      if value == true
        yescount +=1
      else
        yescount = 0
      end

      if i % 5 == 0 
        if yescount == 5 
            puts "Bingo"
            @lastnumber = @calledarray.last
            return true
            # this did it 
        else
          yescount = 0
        end
      
      end
        

    end
      
  end

 
    

   

  #   truecount = 0 
  #   i = 0
  #   bingocardnumbers.each_with_index do |x, i|
  #       if @solvabledictionary[x] == true && truecount <= 5 
  #         truecount += 1
  #         @lastnumber = x
  #         i+=1
  #         # why are we setting the last number here? I guess the assumption is that we will count the last one? 
  #       elsif truecount == 5 && i == 5
  #         puts "I win!"
  #         # okay so the issue here is that we are not counting numbers five at a time
  #         puts "I won after #{index + 1} numbers" 
  #         #the reason this doesn't work is that we're taking an index in, this is going to be an issue, the idea with it is that we take in the number of numbers we have called....
  #         @turnstowin = i
  #         return true
  #       elsif truecount != 5 && i == 5
  #         truecount = 0 
  #         i = 0
  #         puts "at the end of the row, the x value is #{x} and the i value is #{i}"
  #       else
  #         i+=1
  #       end
  
  # end


  def calculatescore
    sum = 0
    # binding.pry
    puts "Before delete the solvabledictionary is #{@solvabledictionary}"
    @solvabledictionary.delete_if { |key, value| value.to_s.match('true') }
    # Okay so that is where we delete the truth values
    # why am I doing that? 
    @solvabledictionary.each_key {|key| sum+= key}
    # winningnumberindex = @calledarray.count() - 2
    # this is wrong as heck 
    # no seriously what the heck is that

    # puts "Winning Number is #{@calledarray[winningnumberindex]}"
    @sum = sum * @calledarray.last
  end

  def execute
    mark_it(@inputnumbers)
    calculatescore
  end


end
# numbers = [7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1]


# Bingo
# inputnumbers = 
#last number does not actually matter 
#card 1 wins 2 after, on 16 
#Winning number does matter, and this thinks they