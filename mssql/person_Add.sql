insert into dbo.Persons
(
	id,
	household_head,
	first_name,
	last_name,
	father_name,
	sex,
	birth_date,
	birth_place,
	identification_data,
	nationality,
	activity_id
)
values
(
	?,
	0,
	?,
	?,
	?,
	substring(?, 1, 1),
	?,
	?,
	?,
	?,
	?
)