select c.*, s.adsh accession_number, si.office, si.industry_title
from companies c
       inner join submissions s on c.cik = s.cik
       left join sics si on s.sic = si.code
where c.symbol = 'AAPL';


select distinct s.form from submissions s


-- select * from submissions where form = '10-Q/A'


select * from submissions where cik = '1592782';