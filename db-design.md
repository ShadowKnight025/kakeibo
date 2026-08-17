
* Expense
  ID
  Title
  DollarAmount - USD
  CategoryID
  Date
  

* Categories
  (UserID, CategoryID, List of Expenses, Monthly_Budget, Budget_Spent, Budget_Remaining)
  
  - Essentials
  - NonEssentials
  - Cultural
  - Unexpected

* Users
  UserID - UUID4
  username
  password(BCRYPT)

--------------------------------------------------
 Logic
 _____
 
 Every Expense adds to the Budget_Spent and Subtracts from Budget_Remaining Totals for the Month per Category.