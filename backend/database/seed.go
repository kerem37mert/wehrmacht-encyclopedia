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
			"Erwin Johannes Eugen Rommel, popularly known as the Desert Fox, was a German field marshal during World War II. He earned fame for his leadership of the Afrika Korps and his tactical brilliance in mobile warfare. Despite his military prowess, he was implicated in the 20 July plot against Hitler and forced to commit suicide. He is remembered as one of the most skilled commanders of the war and was respected by both allies and enemies.",
			"/images/rommel.jpg",
			"Battle of France, North African Campaign, Battle of El Alamein, Battle of Gazala, Siege of Tobruk",
		},
		{
			"Heinz Guderian", "Heer", "Generaloberst",
			"1888-06-17", "1954-05-14",
			"Heinz Wilhelm Guderian was a German general during World War II who played a central role in the development of the panzer division concept and the strategy of Blitzkrieg. Known as the 'father of the Panzer troops', he revolutionized armored warfare with his innovative tactics. He commanded panzer armies during the invasions of Poland, France, and the Soviet Union. After the war, he wrote influential memoirs on armored warfare.",
			"/images/guderian.jpg",
			"Invasion of Poland, Battle of France, Operation Barbarossa, Battle of Smolensk, Drive to Moscow",
		},
		{
			"Erich von Manstein", "Heer", "Generalfeldmarschall",
			"1887-11-24", "1973-06-09",
			"Erich von Manstein was a German commander of the Wehrmacht and was held in high regard as one of the Wehrmacht's best military strategists. He was the architect of the successful Manstein Plan that led to the Fall of France in 1940. On the Eastern Front, he commanded the 11th Army during the siege of Sevastopol and later led Army Group South. His operational brilliance was demonstrated in the Third Battle of Kharkov. He was dismissed by Hitler in 1944 for strategic disagreements.",
			"/images/manstein.jpg",
			"Invasion of Poland, Fall of France, Siege of Sevastopol, Battle of Stalingrad, Battle of Kursk, Third Battle of Kharkov",
		},
		{
			"Karl Dönitz", "Kriegsmarine", "Großadmiral",
			"1891-09-16", "1980-12-24",
			"Karl Dönitz was a German admiral who commanded the German Navy during World War II and briefly succeeded Hitler as head of state. He was the architect of Germany's U-boat warfare strategy and developed the 'wolfpack' tactics that devastated Allied shipping. As Commander-in-Chief of the Kriegsmarine from 1943, he oversaw naval operations until Germany's surrender. He was convicted at Nuremberg and served 10 years in prison.",
			"/images/donitz.jpg",
			"Battle of the Atlantic, U-boat Campaign, Operation Regenbogen",
		},
		{
			"Erich Raeder", "Kriegsmarine", "Großadmiral",
			"1876-04-24", "1960-11-06",
			"Erich Johann Albert Raeder was a German admiral who played a major role in the naval history of World War II. He served as Commander-in-Chief of the Kriegsmarine from 1928 to 1943, overseeing the rebuilding of the German Navy. He planned and executed the invasions of Norway and Denmark. His disagreements with Hitler over naval strategy led to his resignation in 1943. He was sentenced to life imprisonment at Nuremberg but released in 1955.",
			"/images/raeder.jpg",
			"Norwegian Campaign, Operation Weserübung, Battle of the Denmark Strait",
		},
		{
			"Hermann Göring", "Luftwaffe", "Reichsmarschall",
			"1893-01-12", "1946-10-15",
			"Hermann Wilhelm Göring was a German political and military leader and Commander-in-Chief of the Luftwaffe. A World War I flying ace, he became Hitler's designated successor and held multiple high positions in the Nazi regime. He oversaw the Luftwaffe's operations during the Battle of Britain and the Eastern Front. His influence waned after repeated Luftwaffe failures. He was convicted at Nuremberg and committed suicide before his execution.",
			"/images/goring.jpg",
			"Battle of Britain, Operation Barbarossa, Battle of Stalingrad",
		},
		{
			"Albert Kesselring", "Luftwaffe", "Generalfeldmarschall",
			"1885-11-30", "1960-07-16",
			"Albert Kesselring was a German Luftwaffe Generalfeldmarschall during World War II. Known as 'Smiling Albert' for his optimistic demeanor, he commanded Luftwaffe forces during the Battle of Britain and later became Commander-in-Chief South, defending Italy with great skill. His defensive operations in Italy from 1943-1945 demonstrated exceptional tactical ability. He was convicted of war crimes but released in 1952.",
			"/images/kesselring.jpg",
			"Battle of Britain, Mediterranean Theater, Italian Campaign, Battle of Monte Cassino, Gothic Line",
		},
		{
			"Hugo Sperrle", "Luftwaffe", "Generalfeldmarschall",
			"1885-02-07", "1953-04-02",
			"Wilhelm Hugo Sperrle was a German military aviator in World War I and a Generalfeldmarschall in the Luftwaffe during World War II. He commanded the Condor Legion during the Spanish Civil War, gaining valuable combat experience. He led Luftflotte 3 during the Battle of France and the Battle of Britain. Known for his brutal tactics, he was tried for war crimes but acquitted in 1948.",
			"/images/sperrle.jpg",
			"Spanish Civil War, Battle of France, Battle of Britain, Bombing of Rotterdam",
		},
		{
			"Friedrich Paulus", "Heer", "Generalfeldmarschall",
			"1890-09-23", "1957-02-01",
			"Friedrich Wilhelm Ernst Paulus was a German field marshal best known for commanding the 6th Army during the Battle of Stalingrad. He was a skilled staff officer who helped plan Operation Barbarossa. Promoted to field marshal during the Stalingrad battle, he surrendered despite Hitler's orders, becoming the first German field marshal to be captured. After the war, he testified against other German officers at Nuremberg and lived in East Germany.",
			"/images/paulus.jpg",
			"Operation Barbarossa, Battle of Stalingrad, Case Blue",
		},
		{
			"Gerd von Rundstedt", "Heer", "Generalfeldmarschall",
			"1875-12-12", "1953-02-24",
			"Karl Rudolf Gerd von Rundstedt was one of Germany's most senior military commanders during World War II. He commanded Army Group South during the invasion of Poland and Army Group A during the Battle of France. On the Eastern Front, he led Army Group South during Operation Barbarossa. He was Commander-in-Chief West during the D-Day landings. Known for his traditional military professionalism, he was repeatedly dismissed and recalled by Hitler.",
			"/images/rundstedt.jpg",
			"Invasion of Poland, Battle of France, Operation Barbarossa, Battle of Kiev, D-Day, Battle of the Bulge",
		},
		{
			"Wilhelm Keitel", "Heer", "Generalfeldmarschall",
			"1882-09-22", "1946-10-16",
			"Wilhelm Bodewin Johann Gustav Keitel was a German field marshal who served as Chief of the Oberkommando der Wehrmacht (OKW) throughout World War II. He was known as 'Lakeitel' (lackey) by other officers for his servile attitude toward Hitler. He signed numerous criminal orders including the Commando Order and the Commissar Order. He signed Germany's unconditional surrender in Berlin. He was convicted at Nuremberg and executed for war crimes.",
			"/images/keitel.jpg",
			"Invasion of Poland, Battle of France, Operation Barbarossa, German Surrender",
		},
		{
			"Alfred Jodl", "Heer", "Generaloberst",
			"1890-05-10", "1946-10-16",
			"Alfred Josef Ferdinand Jodl was a German Generaloberst who served as Chief of the Operations Staff of the Oberkommando der Wehrmacht throughout World War II. He was one of Hitler's closest military advisors and was present at most major conferences. He signed the unconditional surrender of German forces at Reims. He was convicted of war crimes at Nuremberg and executed, though his conviction was later overturned by a German denazification court.",
			"/images/jodl.jpg",
			"Invasion of Poland, Battle of France, Operation Barbarossa, Battle of Stalingrad, German Surrender",
		},
		{
			"Walter Model", "Heer", "Generalfeldmarschall",
			"1891-01-24", "1945-04-21",
			"Otto Moritz Walter Model was a German field marshal known as the 'Führer's Fireman' for his skill in defensive operations. He excelled at stabilizing collapsing fronts and conducting fighting withdrawals. He commanded various armies and army groups on the Eastern Front, demonstrating exceptional tactical skill. He led German forces during Operation Market Garden and the Battle of the Bulge. He committed suicide in the Ruhr Pocket rather than surrender.",
			"/images/model.jpg",
			"Battle of Rzhev, Battle of Kursk, Operation Bagration, Operation Market Garden, Battle of the Bulge, Ruhr Pocket",
		},
		{
			"Paul Hausser", "Waffen-SS", "SS-Oberst-Gruppenführer",
			"1880-10-07", "1972-12-21",
			"Paul Hausser was one of the most prominent Waffen-SS commanders during World War II. A former Reichswehr officer, he helped establish the SS-Verfügungstruppe and trained the first SS divisions. He commanded the II SS Panzer Corps during the Battle of Kursk and later the 7th Army in Normandy. Despite losing an eye in combat, he continued to command until the end of the war. He was never charged with war crimes.",
			"/images/hausser.jpg",
			"Battle of Kharkov, Battle of Kursk, Battle of Normandy, Battle of the Falaise Pocket",
		},
		{
			"Sepp Dietrich", "Waffen-SS", "SS-Oberst-Gruppenführer",
			"1892-05-28", "1966-04-21",
			"Josef 'Sepp' Dietrich was a German SS commander who led the 1st SS Panzer Division Leibstandarte SS Adolf Hitler and later the 6th Panzer Army. A personal favorite of Hitler, he rose through the ranks despite limited formal military education. He commanded forces during the Battle of the Bulge and the final defense of Germany. He was convicted of war crimes related to the Malmedy massacre and served time in prison.",
			"/images/dietrich.jpg",
			"Battle of France, Operation Barbarossa, Battle of Kharkov, Battle of Normandy, Battle of the Bulge",
		},
		{
			"Kurt Student", "Luftwaffe", "Generaloberst",
			"1890-05-12", "1978-07-01",
			"Kurt Arthur Benno Student was a German Luftwaffe general who pioneered German airborne forces. He commanded the Fallschirmjäger (paratroopers) and planned the airborne assaults on Belgium, the Netherlands, and Crete. The costly victory at Crete led Hitler to abandon large-scale airborne operations. He later commanded ground forces in the Netherlands and northern Germany. He was never convicted of war crimes despite controversial operations.",
			"/images/student.jpg",
			"Battle of Fort Eben-Emael, Battle of Crete, Operation Market Garden, Battle of the Bulge",
		},
		{
			"Eduard Dietl", "Heer", "Generaloberst",
			"1890-07-21", "1944-06-23",
			"Eduard Dietl was a German general known for his leadership in mountain warfare and Arctic operations. He commanded the 3rd Mountain Division during the invasion of Norway, capturing Narvik after fierce fighting. He later commanded the 20th Mountain Army in Finland and northern Russia. Known as the 'Hero of Narvik', he was highly decorated and respected. He died in a plane crash in Austria in 1944.",
			"/images/dietl.jpg",
			"Battle of Narvik, Operation Silver Fox, Arctic Front Operations",
		},
		{
			"Hermann Hoth", "Heer", "Generaloberst",
			"1890-04-12", "1971-01-25",
			"Hermann Hoth was a German panzer commander who led armored forces during the invasions of Poland, France, and the Soviet Union. He commanded the 3rd Panzer Group during Operation Barbarossa and later the 4th Panzer Army at Stalingrad and Kursk. Known for his aggressive tactics and operational skill, he was one of Germany's most capable panzer commanders. He was convicted of war crimes and served 15 years in prison.",
			"/images/hoth.jpg",
			"Invasion of Poland, Battle of France, Operation Barbarossa, Battle of Moscow, Battle of Stalingrad, Battle of Kursk",
		},
		{
			"Ewald von Kleist", "Heer", "Generalfeldmarschall",
			"1881-08-08", "1954-11-13",
			"Paul Ludwig Ewald von Kleist was a German field marshal who commanded panzer armies and army groups during World War II. He led the breakthrough at Sedan during the Battle of France and commanded the 1st Panzer Army during Operation Barbarossa. He later commanded Army Group A in the Caucasus and Army Group South in Ukraine. He was extradited to the Soviet Union after the war and died in captivity.",
			"/images/kleist.jpg",
			"Battle of France, Operation Barbarossa, Battle of Kiev, Case Blue, Battle of the Caucasus",
		},
		{
			"Fedor von Bock", "Heer", "Generalfeldmarschall",
			"1880-12-03", "1945-05-04",
			"Fedor von Bock was a German field marshal who commanded army groups during the invasions of Poland, France, and the Soviet Union. He led Army Group Center during Operation Barbarossa, advancing to the gates of Moscow. Known for his aggressive leadership style, he clashed with Hitler over strategy. He was dismissed in 1942 and briefly recalled in 1943. He was killed by a British fighter-bomber attack near Hamburg in 1945.",
			"/images/bock.jpg",
			"Invasion of Poland, Battle of France, Operation Barbarossa, Battle of Moscow, Case Blue",
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
		{"Blitzkrieg", "Lightning war - a military tactic designed to create disorganization among enemy forces through the use of mobile forces and locally concentrated firepower. Combined arms tactics using tanks, motorized infantry, artillery, and air power to achieve rapid breakthrough and deep penetration.", "Tactics"},
		{"Panzer", "German word for 'armor' or 'tank', commonly used to refer to armored fighting vehicles. Panzer divisions were the backbone of German mobile warfare, combining tanks with mechanized infantry, artillery, and support units.", "Equipment"},
		{"Wehrmacht", "The unified armed forces of Nazi Germany from 1935 to 1945, consisting of the Heer (Army), Kriegsmarine (Navy), and Luftwaffe (Air Force). It replaced the Reichswehr and was disbanded after Germany's defeat.", "Organization"},
		{"Generalfeldmarschall", "Field Marshal - the highest military rank in the Wehrmacht, equivalent to a five-star general. Only 25 officers achieved this rank during World War II. The rank insignia featured crossed batons.", "Ranks"},
		{"Oberkommando der Wehrmacht (OKW)", "High Command of the Armed Forces - the supreme command authority of the Wehrmacht from 1938-1945. Led by Wilhelm Keitel, it coordinated all three service branches and reported directly to Hitler.", "Organization"},
		{"Kesselschlacht", "Cauldron battle or encirclement battle - a tactical maneuver to surround and destroy enemy forces. The Wehrmacht perfected this tactic during the early war years, achieving massive victories at Kiev, Minsk, and Vyazma.", "Tactics"},
		{"Sturmabteilung (SA)", "Storm Detachment - the original paramilitary wing of the Nazi Party. The SA played a key role in Hitler's rise to power but was largely sidelined after the Night of the Long Knives in 1934.", "Organization"},
		{"Schutzstaffel (SS)", "Protection Squadron - a major paramilitary organization under Heinrich Himmler. The Waffen-SS served as elite combat units alongside the Wehrmacht, while other SS branches ran concentration camps and security operations.", "Organization"},
		{"U-Boot", "Abbreviation for Unterseeboot - German submarine. U-boats were the primary weapon in the Battle of the Atlantic, using wolfpack tactics to attack Allied convoys. Type VII and Type IX were the most common variants.", "Equipment"},
		{"Luftwaffe", "Air Force - the aerial warfare branch of the Wehrmacht, commanded by Hermann Göring. At its peak, it was the most powerful air force in the world, featuring advanced aircraft like the Bf 109 and Fw 190 fighters.", "Organization"},
		{"Kriegsmarine", "War Navy - the naval warfare branch of the Wehrmacht. Despite being smaller than the Royal Navy, it achieved notable successes with U-boats and surface raiders like the Bismarck and Graf Spee.", "Organization"},
		{"Heer", "Army - the land warfare branch of the Wehrmacht and its largest component. At its peak in 1943, the Heer had over 6 million personnel organized into hundreds of divisions.", "Organization"},
		{"Stuka", "Sturzkampfflugzeug (dive bomber) - particularly referring to the Junkers Ju 87. Famous for its inverted gull wings and screaming siren, it provided close air support during Blitzkrieg operations but became vulnerable to modern fighters.", "Equipment"},
		{"Flak", "Fliegerabwehrkanone - anti-aircraft gun. The famous 88mm Flak gun was also highly effective against tanks. Flak units protected cities, military installations, and field armies from air attack.", "Equipment"},
		{"Abwehr", "Defense - the German military intelligence service from 1920-1944, headed by Admiral Wilhelm Canaris. It conducted espionage, sabotage, and counterintelligence operations until absorbed by the SS.", "Organization"},
		{"Schwerpunkt", "Point of main effort - a key principle in German military doctrine. Concentrating maximum combat power at the decisive point to achieve breakthrough, rather than spreading forces evenly.", "Tactics"},
		{"Auftragstaktik", "Mission-type tactics - a command philosophy emphasizing initiative and decentralized decision-making. Commanders were given objectives and freedom to determine how to achieve them, enabling flexibility and rapid response.", "Tactics"},
		{"Tiger I", "Panzerkampfwagen VI Tiger - a heavy tank featuring the powerful 88mm gun and thick armor. Introduced in 1942, it dominated Allied armor but was expensive and mechanically complex. Only 1,347 were produced.", "Equipment"},
		{"Panther", "Panzerkampfwagen V Panther - a medium tank designed to counter the Soviet T-34. Featuring sloped armor and a 75mm gun, it's considered one of the best tanks of WWII. Over 6,000 were produced.", "Equipment"},
		{"MG42", "Maschinengewehr 42 - a general-purpose machine gun with a rate of fire up to 1,200 rounds per minute. Known as 'Hitler's Buzzsaw' for its distinctive sound, it remained in service for decades after the war.", "Equipment"},
		{"Fallschirmjäger", "Paratrooper - elite airborne infantry of the Luftwaffe. They conducted successful operations in Norway, Belgium, and Crete, but high casualties at Crete led to their use primarily as elite ground infantry.", "Organization"},
		{"Waffen-SS", "Armed SS - the combat branch of the SS that fought alongside the Wehrmacht. Initially small, it grew to 38 divisions by 1945. Units like the Leibstandarte and Das Reich were considered elite formations.", "Organization"},
		{"Oberkommando des Heeres (OKH)", "High Command of the Army - the supreme command of the German Army. It planned and directed operations on the Eastern Front while OKW handled other theaters.", "Organization"},
		{"Generaloberst", "Colonel General - the second-highest rank in the Wehrmacht, equivalent to a four-star general. Officers like Guderian and Jodl held this rank.", "Ranks"},
		{"Großadmiral", "Grand Admiral - the highest rank in the Kriegsmarine, equivalent to Field Marshal. Only Erich Raeder and Karl Dönitz held this rank during WWII.", "Ranks"},
		{"Reichsmarschall", "Marshal of the Reich - a unique rank created for Hermann Göring in 1940, making him senior to all other Wehrmacht officers. He was the only person to hold this rank.", "Ranks"},
		{"Panzerfaust", "Armor fist - a disposable anti-tank weapon introduced in 1943. Cheap and effective, over 6 million were produced. It could penetrate 200mm of armor at close range.", "Equipment"},
		{"Nebelwerfer", "Smoke thrower - a series of rocket artillery weapons. The six-barreled 15cm Nebelwerfer 41 was nicknamed 'Screaming Mimi' by Allied troops for its distinctive sound.", "Equipment"},
		{"Enigma", "A cipher machine used to encrypt military communications. The breaking of Enigma codes by Allied cryptanalysts at Bletchley Park was a crucial intelligence advantage.", "Equipment"},
		{"Gestapo", "Geheime Staatspolizei (Secret State Police) - the official secret police of Nazi Germany. It ruthlessly suppressed opposition and was involved in numerous war crimes and crimes against humanity.", "Organization"},
		{"Einsatzgruppen", "Deployment groups - SS paramilitary death squads that followed the Wehrmacht into occupied territories. They were responsible for mass murders of Jews, communists, and other groups deemed enemies.", "Organization"},
		{"Festung", "Fortress - Hitler's strategy of declaring cities as fortresses to be defended to the last man. This led to devastating sieges at Stalingrad, Budapest, Breslau, and other cities.", "Tactics"},
		{"Volkssturm", "People's Storm - a national militia established in 1944 consisting of males aged 16-60 not already in military service. Poorly trained and equipped, they were used in the final defense of Germany.", "Organization"},
		{"V-1", "Vergeltungswaffe 1 (Vengeance Weapon 1) - a pulse-jet powered flying bomb, the first cruise missile. Over 10,000 were launched against Britain and Belgium from 1944-1945.", "Equipment"},
		{"V-2", "Vergeltungswaffe 2 - the world's first long-range guided ballistic missile. Developed by Wernher von Braun, over 3,000 were launched against Allied targets. It was too late to affect the war's outcome.", "Equipment"},
		{"Atlantic Wall", "A system of coastal fortifications built along the western coast of Europe to defend against Allied invasion. Despite massive construction efforts, it failed to prevent the D-Day landings.", "Tactics"},
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
			"The German invasion of France and the Low Countries during World War II, also known as the Fall of France. The Wehrmacht executed the Manstein Plan, breaking through the Ardennes and encircling Allied forces. The campaign demonstrated the effectiveness of Blitzkrieg tactics with combined arms operations. France surrendered after just six weeks of fighting, shocking the world.",
			"Germany, France, United Kingdom, Belgium, Netherlands, Luxembourg",
			"Decisive German victory, Fall of France, British evacuation at Dunkirk",
		},
		{
			"Operation Barbarossa", "1941-06-22 to 1941-12-05", "Eastern Front, Soviet Union",
			"The code name for Nazi Germany's invasion of the Soviet Union during World War II, the largest military operation in history. Over 3 million Axis troops attacked along an 1,800-mile front. Initial German advances were spectacular, capturing millions of Soviet soldiers. However, fierce Soviet resistance, vast distances, and the onset of winter halted the offensive before Moscow. This marked the beginning of Germany's eventual defeat.",
			"Germany and Axis powers, Soviet Union",
			"Strategic German failure, Soviet defensive success, turning point of the war",
		},
		{
			"Battle of Stalingrad", "1942-08-23 to 1943-02-02", "Stalingrad, Soviet Union",
			"A major battle on the Eastern Front where Nazi Germany fought the Soviet Union for control of Stalingrad. The battle devolved into brutal house-to-house fighting with massive casualties on both sides. The Soviet counteroffensive Operation Uranus encircled the German 6th Army. Despite Hitler's orders to fight to the death, Field Marshal Paulus surrendered with 91,000 survivors. It was the bloodiest battle in history with nearly 2 million casualties.",
			"Germany and Axis powers, Soviet Union",
			"Decisive Soviet victory, destruction of 6th Army, major turning point of the war",
		},
		{
			"Battle of Britain", "1940-07-10 to 1940-10-31", "United Kingdom airspace",
			"A military campaign when the Royal Air Force defended the United Kingdom against the Luftwaffe's air offensive. Germany sought to gain air superiority as a prelude to invasion (Operation Sea Lion). Despite initial success against RAF airfields, the Luftwaffe shifted to bombing cities (the Blitz). The RAF's victory, aided by radar and the Hurricane/Spitfire fighters, forced Hitler to postpone the invasion indefinitely.",
			"United Kingdom, Germany",
			"British victory, first major defeat for Nazi Germany, invasion of Britain prevented",
		},
		{
			"Battle of El Alamein", "1942-10-23 to 1942-11-11", "El Alamein, Egypt",
			"A battle of the Western Desert Campaign during the Second World War. British forces under Montgomery attacked Rommel's Afrika Korps in a carefully planned offensive. After 12 days of intense fighting, the Axis forces were forced to retreat, ending German hopes of capturing the Suez Canal and Middle Eastern oil fields. Churchill said: 'Before Alamein we never had a victory. After Alamein we never had a defeat.'",
			"British Empire and Allies, Germany and Italy",
			"Allied victory, turning point in North Africa, beginning of Axis retreat",
		},
		{
			"Battle of Kursk", "1943-07-05 to 1943-08-23", "Kursk, Soviet Union",
			"A major World War II Eastern Front battle between German and Soviet forces, the largest tank battle in history. Germany launched Operation Citadel to eliminate the Kursk salient, but Soviet intelligence knew the plans. Massive defensive preparations and counterattacks defeated the German offensive. The battle involved over 6,000 tanks and 2 million troops. After Kursk, Germany never regained the strategic initiative on the Eastern Front.",
			"Germany, Soviet Union",
			"Decisive Soviet victory, end of German offensive capability in the East",
		},
		{
			"Battle of Dunkirk", "1940-05-26 to 1940-06-04", "Dunkirk, France",
			"The evacuation of Allied soldiers from the beaches and harbor of Dunkirk during the Battle of France. Surrounded by German forces, over 338,000 British and French troops were evacuated in Operation Dynamo using military and civilian vessels. Hitler's controversial halt order allowed the evacuation to succeed. While a tactical defeat, the evacuation saved the British Army to fight another day and became a symbol of British resilience.",
			"United Kingdom, France, Germany",
			"Strategic Allied success, 338,000 troops evacuated, but France lost",
		},
		{
			"Battle of Moscow", "1941-10-02 to 1942-01-07", "Moscow, Soviet Union",
			"The German offensive to capture Moscow (Operation Typhoon) and the subsequent Soviet counteroffensive. German forces came within sight of the Kremlin but were halted by fierce Soviet resistance, harsh winter conditions, and overextended supply lines. The Soviet counterattack drove German forces back 100-250 km, the first major German defeat of the war. It shattered the myth of Wehrmacht invincibility.",
			"Germany and Axis powers, Soviet Union",
			"Soviet victory, first major German defeat, Moscow saved",
		},
		{
			"Battle of the Bulge", "1944-12-16 to 1945-01-25", "Ardennes, Belgium and Luxembourg",
			"Germany's last major offensive campaign on the Western Front. Hitler launched a surprise attack through the Ardennes to split Allied forces and capture Antwerp. Initial German success created a 'bulge' in Allied lines. However, American resistance at Bastogne, clearing weather allowing air support, and fuel shortages doomed the offensive. The battle depleted Germany's strategic reserves, hastening the end of the war.",
			"United States, United Kingdom, Germany",
			"Allied victory, German strategic reserves exhausted, path to Germany opened",
		},
		{
			"Battle of Crete", "1941-05-20 to 1941-06-01", "Crete, Greece",
			"The first major airborne invasion in military history. German Fallschirmjäger and mountain troops attacked the island defended by British, Greek, and Commonwealth forces. Despite heavy casualties (over 50% in some units), Germany captured the island. However, losses were so severe that Hitler never authorized another large-scale airborne operation. The battle demonstrated both the potential and risks of airborne warfare.",
			"Germany, United Kingdom, Greece, New Zealand, Australia",
			"German victory but at heavy cost, no more major German airborne operations",
		},
		{
			"Siege of Leningrad", "1941-09-08 to 1944-01-27", "Leningrad, Soviet Union",
			"One of the longest and most destructive sieges in history, lasting 872 days. German and Finnish forces encircled Leningrad, subjecting it to constant bombardment and starvation. Over 1 million civilians died, mostly from starvation. Despite horrific conditions, the city never surrendered. The siege was finally broken by Soviet offensives in 1944. The siege became a symbol of Soviet resistance and suffering.",
			"Germany, Finland, Soviet Union",
			"Soviet victory, city held, massive civilian casualties",
		},
		{
			"Operation Market Garden", "1944-09-17 to 1944-09-25", "Netherlands",
			"An Allied airborne operation to capture bridges across the Rhine and end the war by Christmas 1944. British, American, and Polish airborne forces seized bridges while ground forces advanced. The operation failed when British paratroopers at Arnhem were defeated by unexpected SS Panzer divisions. The bridge 'too far' remained in German hands. The failure prolonged the war and led to the Dutch 'Hunger Winter.'",
			"United Kingdom, United States, Poland, Germany",
			"German victory, Allied advance halted, war prolonged",
		},
		{
			"Siege of Sevastopol", "1941-10-30 to 1942-07-04", "Sevastopol, Crimea",
			"A prolonged siege of the Soviet naval base by German and Romanian forces. The fortress was defended by Soviet naval infantry and army units with strong fortifications. Germany employed massive artillery including the 800mm Schwerer Gustav railway gun. After eight months of intense fighting and three major assaults, the city fell. The siege demonstrated the strength of prepared defenses and the cost of reducing fortified positions.",
			"Germany, Romania, Soviet Union",
			"German victory, Crimea secured, but at high cost in time and casualties",
		},
		{
			"Third Battle of Kharkov", "1943-02-19 to 1943-03-15", "Kharkov, Ukraine",
			"A German counteroffensive following the disaster at Stalingrad. Manstein's Army Group South executed a brilliant mobile defense and counterattack, recapturing Kharkov and destroying several Soviet armies. The battle demonstrated German tactical superiority even in retreat. It was Germany's last major victory on the Eastern Front and stabilized the front before the Battle of Kursk.",
			"Germany, Soviet Union",
			"German tactical victory, Kharkov recaptured, front stabilized",
		},
		{
			"Battle of Monte Cassino", "1944-01-17 to 1944-05-18", "Monte Cassino, Italy",
			"A series of four assaults by Allied forces to break through the German Winter Line and advance on Rome. The historic Monte Cassino monastery dominated the battlefield and was controversially bombed. German Fallschirmjäger defended the ruins with great skill. After four months of brutal fighting, Polish forces finally captured the monastery. The battle opened the road to Rome but at a heavy cost.",
			"Allies (US, UK, Poland, France, India, New Zealand), Germany",
			"Allied victory, Winter Line broken, road to Rome opened",
		},
		{
			"Battle of Normandy", "1944-06-06 to 1944-08-30", "Normandy, France",
			"The Allied invasion of German-occupied France, beginning with D-Day landings on June 6, 1944. Over 150,000 Allied troops landed on five beaches, establishing a foothold despite fierce German resistance. The subsequent campaign saw intense fighting in the bocage countryside. The German army was encircled and destroyed in the Falaise Pocket. The liberation of France had begun.",
			"Allies (US, UK, Canada, France), Germany",
			"Decisive Allied victory, liberation of France begun, Germany's western front collapsed",
		},
		{
			"Battle of Kiev", "1941-08-23 to 1941-09-26", "Kiev, Ukraine",
			"The largest encirclement battle in military history. German forces under Rundstedt and Guderian surrounded four Soviet armies defending Kiev. Despite Stalin's orders to hold the city, over 600,000 Soviet soldiers were captured. The victory opened the path to the Donbas industrial region but delayed the advance on Moscow, possibly saving the Soviet capital.",
			"Germany, Soviet Union",
			"Decisive German victory, largest encirclement in history, 600,000+ captured",
		},
		{
			"Battle of the Atlantic", "1939-09-03 to 1945-05-08", "Atlantic Ocean",
			"The longest continuous military campaign of World War II. German U-boats attempted to cut Britain's supply lines while Allied navies and air forces hunted the submarines. The introduction of convoys, escort carriers, long-range aircraft, and breaking the Enigma code gradually turned the tide. By 1943, Allied victory was assured. Over 3,500 merchant ships and 783 U-boats were lost.",
			"Allies (UK, US, Canada), Germany",
			"Allied victory, sea lanes secured, Germany's submarine force destroyed",
		},
		{
			"Invasion of Poland", "1939-09-01 to 1939-10-06", "Poland",
			"The German invasion that began World War II. Wehrmacht forces attacked from the west while Soviet forces invaded from the east on September 17. The campaign demonstrated Blitzkrieg tactics with rapid armored thrusts and close air support. Despite brave Polish resistance, the country was overrun in five weeks. The invasion led to Britain and France declaring war on Germany.",
			"Germany, Soviet Union, Poland",
			"German and Soviet victory, Poland partitioned, World War II begins",
		},
		{
			"Battle of Berlin", "1945-04-16 to 1945-05-02", "Berlin, Germany",
			"The final major offensive of the European theater. Over 2.5 million Soviet troops attacked Berlin defended by Wehrmacht, SS, Hitler Youth, and Volkssturm. The battle was characterized by brutal urban warfare with fighting for every building. Hitler committed suicide on April 30. The city surrendered on May 2, effectively ending the war in Europe. Over 1 million casualties made it one of the bloodiest battles.",
			"Soviet Union, Poland, Germany",
			"Soviet victory, fall of Berlin, end of Nazi Germany, Hitler's suicide",
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
		{1, "The best form of welfare for the troops is first-rate training.", "On military training", "1941"},
		{2, "The engine of the panzer is a weapon just as the main gun.", "On the mobility of armored warfare", "1937"},
		{2, "If the tanks succeed, then victory follows.", "On the decisive role of armor", "1940"},
		{2, "You hit somebody with your fist and not with your fingers spread.", "On concentration of force", "1939"},
		{3, "The man who gets to the battlefield first with the most men wins.", "On the importance of initiative", "1941"},
		{3, "War is the domain of danger, therefore courage is the soldier's first requirement.", "On military courage", "1942"},
		{3, "The commander must try, above all, to establish personal and comradely contact with his men.", "On leadership", "1943"},
		{4, "The submarine is the capital ship of the future.", "On naval warfare strategy", "1939"},
		{4, "The best weapon against the enemy submarine is another submarine.", "On U-boat tactics", "1940"},
		{5, "A ship is always referred to as 'she' because it costs so much to keep one in paint and powder.", "On naval humor", "1938"},
		{6, "Guns will make us powerful; butter will only make us fat.", "On military priorities", "1936"},
		{6, "Shoot first and ask questions later.", "On decisive action", "1940"},
		{7, "Flexibility is the key to stability.", "On military adaptation", "1943"},
		{7, "A good plan violently executed now is better than a perfect plan executed next week.", "On the importance of timing", "1944"},
		{9, "I have no intention of shooting myself for this Bohemian corporal.", "On his decision to surrender at Stalingrad", "1943"},
		{10, "The German soldier has impressed the world; however the German officer blundered into the war unprepared.", "On German military leadership", "1945"},
		{10, "I do not believe in miracles, but in hard work and careful planning.", "On military professionalism", "1942"},
		{13, "There are no desperate situations, there are only desperate people.", "On maintaining morale in difficult circumstances", "1944"},
		{13, "Attack, attack, attack! Never give the enemy time to recover.", "On offensive operations", "1944"},
		{16, "He who holds Crete holds the Eastern Mediterranean.", "On the strategic importance of Crete", "1941"},
		{16, "The parachute is the weapon of the future.", "On airborne warfare", "1940"},
		{17, "In the north, one must be hard as the ice and cold as the snow.", "On Arctic warfare", "1942"},
		{18, "The tank is the modern cavalry.", "On armored warfare doctrine", "1941"},
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
