declare @city_1_id int = ?
declare @city_2_id int = ?
insert into dbo.Tiers
(
	creation_date,
	client_type_code,
	loan_cycle,
	active,
	bad_client,
	branch_id,
	home_phone,
	personal_phone,
	e_mail,
	district_id,
	city,
	address,
	zipCode,
	secondary_district_id,
	secondary_city,
	secondary_address,
	secondary_zipCode
)
select
	getdate(),
	'I',
	0,
	0,
	0,
	?,
	?,
	?,
	?,
	(select district_id from dbo.City where id = @city_1_id),
	(select name from dbo.City where id = @city_1_id),
	?,
	?,
	(select district_id from dbo.City where id = @city_2_id),
	(select name from dbo.City where id = @city_2_id),
	?,
	?
