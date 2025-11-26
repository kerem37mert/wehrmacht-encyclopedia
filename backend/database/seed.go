package database

import (
	"log"
)

// SeedData populates the database with sample data
func SeedData() error {
	// Check if data already exists
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM generals").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	// Seed Generals
	generals := []struct {
		name, branch, rank, birth, death, bio, photo, battles string
	}{
		{
			"Erwin Rommel", "Heer", "Generalfeldmarschall",
			"1891-11-15", "1944-10-14",
			"Erwin Johannes Eugen Rommel, popularly known as the Desert Fox, was a German field marshal during World War II.",
			"/images/rommel.jpg",
			"Battle of France, North African Campaign, Battle of El Alamein",
		},
		{
			"Heinz Guderian", "Heer", "Generaloberst",
			"1888-06-17", "1954-05-14",
			"Heinz Wilhelm Guderian was a German general during World War II who played a central role in the development of the panzer division concept.",
			"/images/guderian.jpg",
			"Invasion of Poland, Battle of France, Operation Barbarossa",
		},
		{
			"Erich von Manstein", "Heer", "Generalfeldmarschall",
			"1887-11-24", "1973-06-09",
			"Erich von Manstein was a German commander of the Wehrmacht and was held in high regard as one of the Wehrmacht's best military strategists.",
			"/images/manstein.jpg",
			"Invasion of Poland, Fall of France, Battle of Stalingrad, Battle of Kursk",
		},
		{
			"Karl Dönitz", "Kriegsmarine", "Großadmiral",
			"1891-09-16", "1980-12-24",
			"Karl Dönitz was a German admiral who commanded the German Navy during World War II.",
			"/images/donitz.jpg",
			"Battle of the Atlantic, U-boat Campaign",
		},
		{
			"Erich Raeder", "Kriegsmarine", "Großadmiral",
			"1876-04-24", "1960-11-06",
			"Erich Johann Albert Raeder was a German admiral who played a major role in the naval history of World War II.",
			"/images/raeder.jpg",
			"Norwegian Campaign, Operation Weserübung",
		},
		{
			"Hermann Göring", "Luftwaffe", "Reichsmarschall",
			"1893-01-12", "1946-10-15",
			"Hermann Wilhelm Göring was a German political and military leader and Commander-in-Chief of the Luftwaffe.",
			"/images/goring.jpg",
			"Battle of Britain, Operation Barbarossa",
		},
		{
			"Albert Kesselring", "Luftwaffe", "Generalfeldmarschall",
			"1885-11-30", "1960-07-16",
			"Albert Kesselring was a German Luftwaffe Generalfeldmarschall during World War II.",
			"/images/kesselring.jpg",
			"Battle of Britain, Mediterranean Theater, Italian Campaign",
		},
		{
			"Hugo Sperrle", "Luftwaffe", "Generalfeldmarschall",
			"1885-02-07", "1953-04-02",
			"Wilhelm Hugo Sperrle was a German military aviator in World War I and a Generalfeldmarschall in the Luftwaffe during World War II.",
			"/images/sperrle.jpg",
			"Spanish Civil War, Battle of France, Battle of Britain",
		},
	}

	for _, g := range generals {
		_, err := DB.Exec(`INSERT INTO generals (name, branch, rank, birth_date, death_date, biography, photo_url, notable_battles)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			g.name, g.branch, g.rank, g.birth, g.death, g.bio, g.photo, g.battles)
		if err != nil {
			return err
		}
	}

	// Seed Terms
	terms := []struct {
		term, definition, category string
	}{
		{"Blitzkrieg", "Lightning war - a military tactic designed to create disorganization among enemy forces.", "Tactics"},
		{"Panzer", "German word for 'armor' or 'tank', commonly used to refer to armored fighting vehicles.", "Equipment"},
		{"Wehrmacht", "The unified armed forces of Nazi Germany from 1935 to 1945.", "Organization"},
		{"Generalfeldmarschall", "Field Marshal - the highest military rank in the Wehrmacht.", "Ranks"},
		{"Oberkommando der Wehrmacht (OKW)", "High Command of the Armed Forces.", "Organization"},
		{"Kesselschlacht", "Cauldron battle or encirclement battle.", "Tactics"},
		{"Sturmabteilung (SA)", "Storm Detachment - the original paramilitary wing of the Nazi Party.", "Organization"},
		{"Schutzstaffel (SS)", "Protection Squadron - a major paramilitary organization.", "Organization"},
		{"U-Boot", "Abbreviation for Unterseeboot - German submarine.", "Equipment"},
		{"Luftwaffe", "Air Force - the aerial warfare branch of the Wehrmacht.", "Organization"},
		{"Kriegsmarine", "War Navy - the naval warfare branch of the Wehrmacht.", "Organization"},
		{"Heer", "Army - the land warfare branch of the Wehrmacht.", "Organization"},
		{"Stuka", "Sturzkampfflugzeug (dive bomber) - particularly referring to the Junkers Ju 87.", "Equipment"},
		{"Flak", "Fliegerabwehrkanone - anti-aircraft gun.", "Equipment"},
		{"Abwehr", "Defense - the German military intelligence service.", "Organization"},
	}

	for _, t := range terms {
		_, err := DB.Exec(`INSERT INTO terms (term, definition, category) VALUES (?, ?, ?)`,
			t.term, t.definition, t.category)
		if err != nil {
			return err
		}
	}

	// Seed Battles
	battles := []struct {
		name, date, location, description, participants, outcome string
	}{
		{
			"Battle of France", "1940-05-10 to 1940-06-25", "France, Low Countries",
			"The German invasion of France and the Low Countries during World War II.",
			"Germany, France, United Kingdom, Belgium, Netherlands",
			"Decisive German victory, Fall of France",
		},
		{
			"Operation Barbarossa", "1941-06-22 to 1941-12-05", "Eastern Front, Soviet Union",
			"The code name for Nazi Germany's invasion of the Soviet Union during World War II.",
			"Germany and Axis powers, Soviet Union",
			"Strategic German failure, Soviet defensive success",
		},
		{
			"Battle of Stalingrad", "1942-08-23 to 1943-02-02", "Stalingrad, Soviet Union",
			"A major battle on the Eastern Front where Nazi Germany fought the Soviet Union for control of Stalingrad.",
			"Germany and Axis powers, Soviet Union",
			"Decisive Soviet victory, turning point of the war",
		},
		{
			"Battle of Britain", "1940-07-10 to 1940-10-31", "United Kingdom airspace",
			"A military campaign when the Royal Air Force defended the United Kingdom against the Luftwaffe.",
			"United Kingdom, Germany",
			"British victory, first major defeat for Nazi Germany",
		},
		{
			"Battle of El Alamein", "1942-10-23 to 1942-11-11", "El Alamein, Egypt",
			"A battle of the Western Desert Campaign during the Second World War.",
			"British Empire and Allies, Germany and Italy",
			"Allied victory, turning point in North Africa",
		},
		{
			"Battle of Kursk", "1943-07-05 to 1943-08-23", "Kursk, Soviet Union",
			"A major World War II Eastern Front battle between German and Soviet forces.",
			"Germany, Soviet Union",
			"Decisive Soviet victory, end of German offensive capability",
		},
	}

	for _, b := range battles {
		_, err := DB.Exec(`INSERT INTO battles (name, date, location, description, participants, outcome)
			VALUES (?, ?, ?, ?, ?, ?)`,
			b.name, b.date, b.location, b.description, b.participants, b.outcome)
		if err != nil {
			return err
		}
	}

	// Seed Quotes
	quotes := []struct {
		generalID        int
		quote, ctx, date string
	}{
		{1, "In a man-to-man fight, the winner is he who has one more round in his magazine.", "On the importance of preparation", "1942"},
		{1, "Don't fight a battle if you don't gain anything by winning.", "On strategic thinking", "1943"},
		{2, "The engine of the panzer is a weapon just as the main gun.", "On the mobility of armored warfare", "1937"},
		{2, "If the tanks succeed, then victory follows.", "On the decisive role of armor", "1940"},
		{3, "The man who gets to the battlefield first with the most men wins.", "On the importance of initiative", "1941"},
		{4, "The submarine is the capital ship of the future.", "On naval warfare strategy", "1939"},
		{6, "Guns will make us powerful; butter will only make us fat.", "On military priorities", "1936"},
		{7, "Flexibility is the key to stability.", "On military adaptation", "1943"},
	}

	for _, q := range quotes {
		_, err := DB.Exec(`INSERT INTO quotes (general_id, quote_text, context, date)
			VALUES (?, ?, ?, ?)`,
			q.generalID, q.quote, q.ctx, q.date)
		if err != nil {
			return err
		}
	}

	log.Println("Database seeded successfully")
	return nil
}
