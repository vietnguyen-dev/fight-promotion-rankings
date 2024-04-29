import { useState } from 'react'

import { iWeightClass } from '../interfaces/weight-classes'
import { iRankings } from '../interfaces/ranking'

interface iAccordianProps {
	weightClasses: iWeightClass[],
	rankings: iRankings[]
}
function Accordian({ weightClasses, rankings }: iAccordianProps) {
	//create an object that controls the state of each weight class div, onclick => show the runner up panel

	return (
		<>
			{weightClasses.map((weightClass) => {
				const ranked = rankings.filter((ranking) => ranking.weight_class_id == weightClass.id)
				const champ = ranked[0]
				return (
					<div key={weightClass.id} id={`weight-class-${weightClass.id}`}>
						<div className="p-6 border-[1px] border-b-gray-300">
							<h3 className="text-sm font-bold text-red-700">{weightClass.name.toUpperCase()}</h3>
							{champ ? <h3 className="text-xl font-bold mb-1">{champ.athlete_name.toUpperCase()}</h3> : <h3 className="mb-1">VACANT</h3>}
							<h4 className="text-xs text-gray-600">CHAMPION</h4>
						</div>
						<div className="px-3 hidden">
							<h5>runner ups</h5>
						</div>
					</div>
				);
			})}
		</>
	);
}

export default Accordian
