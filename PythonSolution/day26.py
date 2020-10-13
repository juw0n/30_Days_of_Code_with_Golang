return_day, return_month, return_year = (int(e) for e in input().strip().split(' '))
due_day, due_month, due_year = (int(e) for e in input().strip().split(' '))

'''
fine = 0
if return_year > due_year:
    fine = 10000
elif return_year == due_year:
    #check month
    if return_month > due_month:
        fine = 500 * (return_month - due_month)
    elif return_month == due_month:
        #check day
        if return_day > due_day:
            fine = 15 *(return_day - due_day)
print(fine)
'''

#OR


fine = 0
if return_year < due_year:
    print(fine)
elif return_year == due_year:
    #check month
    if return_month > due_month:
          fine = 500 * (return_month - due_month)
    elif return_month == due_month:
        #check day
        if return_day > due_day:
            fine = 15 *(return_day - due_day)
print(fine)