import { useState } from 'react'
import Link from 'next/link'

import { iWeightClass } from '../interfaces/weight-classes'
import { iRankings } from '../interfaces/ranking'

interface iAccordianProps {
	weightClasses: iWeightClass[],
	rankings: iRankings[]
}

type WeightClassesExpandedState = Record<number, boolean>
function Accordian({ weightClasses, rankings }: iAccordianProps) {
	// Initialize state with default values (all collapsed)
	const [isExpanded, setIsExpanded] = useState<WeightClassesExpandedState>(() => {
		const initialState: WeightClassesExpandedState = {};
		weightClasses.forEach((wc: iWeightClass) => {
			initialState[wc.id] = true; // All sections start collapsed
		});
		return initialState;
	});

	// Function to toggle the state for a specific weight class
	const toggleExpansion = (id: number) => {
		setIsExpanded((prevState) => ({
			...prevState,
			[id]: !prevState[id], // Toggle the specific section
		}));
	};

	return (
		<>
			{weightClasses.map((weightClass: iWeightClass) => {
				const ranked = rankings.filter((ranking) => ranking.weight_class_id == weightClass.id)
				const champ = ranked[0]
				const runnerUps = ranked.slice(1, ranked.length)
				return (
					<section key={weightClass.id} id={`weight-class-${weightClass.id}`}>
						<div className="p-6 border-[1px] border-b-gray-300 flex justify-between">
							<div>
								<h3 className="text-sm font-bold text-red-700">{weightClass.name.toUpperCase()}</h3>
								{champ ?
									<h3 className="text-xl font-bold mb-1">
										<Link href={`/athletes/${champ.athlete_id}`}>{champ.athlete_name.toUpperCase()}</Link>
									</h3>
									: <h3 className="mb-1">VACANT</h3>}
								<h4 className="text-xs text-gray-600">CHAMPION</h4>
							</div>
							<button onClick={() => toggleExpansion(weightClass.id)}>{isExpanded[weightClass.id] ? <p>up</p> : <p>down</p>}</button>
						</div>
						<div className={` ${isExpanded[weightClass.id] && 'hidden'}`}>
							{runnerUps.map((ranking: iRankings) =>
								<div key={ranking.athlete_id} className="flex px-6 py-4 border-b-gray-300">
									<h4 className="font-semibold text-sm text-gray-500 mr-3">{ranking.ranking}</h4>
									<h4 className="font-bold text-sm">
										<Link href={`/athletes/${ranking.athlete_id}`}>{ranking.athlete_name.toUpperCase()}</Link>
									</h4>
								</div>)}

						</div>
					</section >
				);
			})}
		</>
	);
}

export default Accordian
