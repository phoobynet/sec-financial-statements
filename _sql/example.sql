select c.*, s.adsh accession_number
from companies c
       inner join submissions s on c.cik = s.cik
where c.symbol = 'AAPL'
  and s.form = '10-Q'