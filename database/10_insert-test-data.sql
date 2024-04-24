--promotions
INSERT INTO promotions (pro_code, name, email, how_many_ranked) VALUES ('NWKBL0','Northwest Kickboxing League', 'vietnguyent22@gmail.com', 5); 

--results
INSERT INTO results (res_code, name) VALUES ('TKO', 'Total Knockout');
INSERT INTO results (res_code, name) VALUES ('KOT', 'Knockout');
INSERT INTO results (res_code, name) VALUES ('SUB', 'Submission');
INSERT INTO results (res_code, name) VALUES ('UNA', 'Uniamous Descision');
INSERT INTO results (res_code, name) VALUES ('SPD', 'Split Decision');
INSERT INTO results (res_code, name) VALUES ('DIS', 'Disqualification');
INSERT INTO results (res_code, name) VALUES ('DRA', 'Draw');
INSERT INTO results (res_code, name) VALUES ('NOC', 'No Contest');

--ranking placement
INSERT INTO ranking_placement (ran_code, name) VALUES ('CHA', 'Champion');
INSERT INTO ranking_placement (ran_code, name) VALUES ('FIR', '1.');
INSERT INTO ranking_placement (ran_code, name) VALUES ('SEC', '2.');
INSERT INTO ranking_placement (ran_code, name) VALUES ('THI', '3.');
INSERT INTO ranking_placement (ran_code, name) VALUES ('FOR', '4.');
INSERT INTO ranking_placement (ran_code, name) VALUES ('FIF', '5.');
INSERT INTO ranking_placement (ran_code, name) VALUES ('SIX', '6.');
INSERT INTO ranking_placement (ran_code, name) VALUES ('SEV', '7.');
INSERT INTO ranking_placement (ran_code, name) VALUES ('EIG', '8.');
INSERT INTO ranking_placement (ran_code, name) VALUES ('NIN', '9.');
INSERT INTO ranking_placement (ran_code, name) VALUES ('TEN', '10.');

--weight-classes
INSERT INTO weight_classes (wei_code, weight_range, name, promotion_id) VALUES ('BAN135', '135', 'Bantamweight 135', 1);  
INSERT INTO weight_classes (wei_code, weight_range, name, promotion_id) VALUES ('CAT140', '140', 'Catchweight 140', 1);  
INSERT INTO weight_classes (wei_code, weight_range, name, promotion_id) VALUES ('CAT150', '150', 'Catchweight 150', 1);  
INSERT INTO weight_classes (wei_code, weight_range, name, promotion_id) VALUES ('LIG155', '155', 'Light 155', 1);  
INSERT INTO weight_classes (wei_code, weight_range, name, promotion_id) VALUES ('CAT160', '160', 'Catchweight 160', 1);  
INSERT INTO weight_classes (wei_code, weight_range, name, promotion_id) VALUES ('CAT175', '175', 'Catchweight 175', 1);  
INSERT INTO weight_classes (wei_code, weight_range, name, promotion_id) VALUES ('CAT180', '180', 'Catchweight 180', 1);  
INSERT INTO weight_classes (wei_code, weight_range, name, promotion_id) VALUES ('HEA200', '200', 'Heavyweight 200', 1);  

--athletes
INSERT INTO athletes (first_name, last_name, nickname, gym_name, location, height, reach, promotion_id) VALUES ('Keith', 'Nguyen', 'Kamikaze', 'Tigerstyle Combat Sports', 'Portland, OR', '6" 1"', '73', 1);
INSERT INTO athletes (first_name, last_name, gym_name, location, height, reach, promotion_id) VALUES ('Eli', 'Davis', 'Mercenary Combat Sports', 'Pasco, WA', '6" 1"', '73',1 );

INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Bryan', 'Warner', 1);
INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Hunter', 'Walh', 1);

INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Steven', 'Dopp', 1);
INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Alex', 'Morales', 1);

INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Nathan', 'Newman', 1);
INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Zach', 'Robinson', 1);

INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Axel', 'Hoang', 1);
INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Moath', 'Nassar', 1);


INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Dominick', 'Hall', 1);
INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Isaiah', 'Archiaga', 1);

INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Brandon', 'Hendley', 1);
INSERT INTO athletes (first_name, last_name, promotion_id) VALUES ('Brent', 'Campbell', 1);

--events
INSERT INTO events (eve_code, event_title, promotion_id, event_date) VALUES ('NWKBL7', 'Northwest Kickboxing League: 7', 1, '2023-12-30');

--fights
INSERT INTO fights (athlete_id, athlete_two_id, events_id, weight_class_id, time_ellapsed, rounds_lasted, number_rounds, min_per_round) VALUES (1, 2, 1, 5, '0:17', 1, 3, 3);
INSERT INTO fights (athlete_id, athlete_two_id, events_id, weight_class_id, number_rounds, min_per_round) VALUES (3, 4, 1, 7, 3, 3);
INSERT INTO fights (athlete_id, athlete_two_id, events_id, weight_class_id, number_rounds, min_per_round) VALUES (5, 6, 1, 1, 3, 3);
INSERT INTO fights (athlete_id, athlete_two_id, events_id, weight_class_id, number_rounds, min_per_round) VALUES (7, 8, 1, 3, 3, 3);
INSERT INTO fights (athlete_id, athlete_two_id, events_id, weight_class_id, number_rounds, min_per_round) VALUES (9, 10, 1, 4, 3, 3);
INSERT INTO fights (athlete_id, athlete_two_id, events_id, weight_class_id, number_rounds, min_per_round) VALUES (11, 12, 1, 7, 3, 3);
INSERT INTO fights (athlete_id, athlete_two_id, events_id, weight_class_id, number_rounds, min_per_round) VALUES (13, 14, 1, 2, 3, 3);

--fight results
INSERT INTO fight_results (fight_id, winner_id, loser_id, result_id, event_id) VALUES (1, 2, 1, 2, 1); 
INSERT INTO fight_results (fight_id, winner_id, loser_id, result_id, event_id) VALUES (2, 4, 3, 2, 1); 
INSERT INTO fight_results (fight_id, winner_id, loser_id, result_id, event_id) VALUES (3, 6, 5, 5, 1); 
INSERT INTO fight_results (fight_id, winner_id, loser_id, result_id, event_id) VALUES (4, 7, 8, 2, 1); 
INSERT INTO fight_results (fight_id, winner_id, loser_id, result_id, event_id) VALUES (5, 9, 10, 2, 1); 
INSERT INTO fight_results (fight_id, winner_id, loser_id, result_id, event_id) VALUES (6, 11, 12, 5, 1); 
INSERT INTO fight_results (fight_id, winner_id, loser_id, result_id, event_id) VALUES (7, 13, 14, 1, 1); 

--rankings
INSERT INTO rankings (promotion_id, ranking_placement_id, weight_class_id, athlete_id) VALUES (1, 1, 8, 13);
