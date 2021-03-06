package aol

import (
	"math/rand"
	"time"
)

type RandName struct {
	FirstName string
	LastName  string
}

func NewRandName() RandName {
	rand.Seed(time.Now().UnixNano())

	firstNames := []string{"Ada", "Adela", "Adelaida", "Adelina", "Adina", "Adriana", "Agata", "Aglaia", "Agripina", "Aida", "Alberta", "Albertina", "Alexandra", "Alexandrina", "Alice", "Alida", "Alina", "Alis", "Alma", "Amalia", "Amanda", "Amelia", "Ana", "Anabela", "Anaida", "Anamaria", "Anastasia", "Anca", "Ancuta", "Anda", "Andra", "Andrada", "Andreea", "Anemona", "Aneta", "Angela", "Anghelina", "Anica", "Anisoara", "Antoaneta", "Antonela", "Antonia", "Anuta", "Ariadna", "Ariana", "Arina", "Aristita", "Artemisa", "Astrid", "Atena", "Augustina", "Aura", "Aurelia", "Aureliana", "Aurica", "Aurora", "Axenia", "Beatrice", "Betina", "Bianca", "Blanduzia", "Bogdana", "Brândusa", "Camelia", "Carina", "Carla", "Carmen", "Carmina", "Carolina", "Casandra", "Casiana", "Caterina", "Catinca", "Catrina", "Catrinel", "Catalina", "Cecilia", "Celia", "Cerasela", "Cezara", "Cipriana", "Clara", "Clarisa", "Claudia", "Clementina", "Cleopatra", "Codrina", "Codruta", "Constanta", "Constantina", "Consuela", "Coralia", "Corina", "Cornelia", "Cosmina", "Crenguta", "Crina", "Cristina", "Daciana", "Dafina", "Daiana", "Dalia", "Dana", "Daniela", "Daria", "Dariana", "Delia", "Demetra", "Denisa", "Despina", "Diana", "Dida", "Didina", "Dimitrina", "Dina", "Dochia", "Doina", "Domnica", "Dora", "Doriana", "Dorina", "Dorli", "Draga", "Dumitra", "Dumitrana", "Ecaterina", "Eftimia", "Elena", "Eleonora", "Eliana", "Elisabeta", "Elisaveta", "Eliza", "Elodia", "Elvira", "Emanuela", "Emilia", "Erica", "Estera", "Eufrosina", "Eugenia", "Eusebia", "Eva", "Evanghelina", "Evelina", "Fabia", "Fabiana", "Felicia", "Filofteia", "Fiona", "Flavia", "Floare", "Floarea", "Flora", "Florenta", "Florentina", "Floriana", "Florica", "Florina", "Francesca", "Frusina", "Gabriela", "Geanina", "Gentiana", "Georgeta", "Georgia", "Georgiana", "Geta", "Gherghina", "Gianina", "Gina", "Giorgiana", "Gratiana", "Gratiela", "Henrieta", "Heracleea", "Hortensia", "Iasmina", "Ica", "Ileana", "Ilinca", "Ilona", "Ina", "Ioana", "Ioanina", "Iolanda", "Ionela", "Ionelia", "Iosefina", "Iridenta", "Irina", "Iris", "Isabela", "Iulia", "Iuliana", "Iustina", "Ivona", "Izabela", "Jana", "Janeta", "Janina", "Jasmina", "Jeana", "Julia", "Julieta", "Larisa", "Laura", "Laurentia", "Lavinia", "Lacramioara", "Leana", "Lelia", "Leontina", "Leopoldina", "Letitia", "Lia", "Liana", "Lidia", "Lidia", "Ligia", "Lili", "Liliana", "Lioara", "Livia", "Loredana", "Lorelei", "Lorena", "Luana", "Lucia", "Luciana", "Lucretia", "Ludmila", "Ludovica", "Luiza", "Luminita", "Magdalena", "Maia", "Malvina", "Manuela", "Mara", "Marcela", "Marcheta", "Marga", "Margareta", "Maria", "Mariana", "Maricica", "Marilena", "Marina", "Marinela", "Marioara", "Marta", "Matilda", "Madalina", "Malina", "Marioara", "Mariuca", "Melania", "Melina", "Mihaela", "Milena", "Minodora", "Mioara", "Mirabela", "Mirela", "Mirona", "Miruna", "Mona", "Monalisa", "Monica", "Nadia", "Narcisa", "Natalia", "Natasa", "Nicoleta", "Niculina", "Nidia", "Noemi", "Nora", "Norica", "Oana", "Octavia", "Octaviana", "Ofelia", "Olga", "Olimpia", "Olivia", "Ortansa", "Otilia", "Ozana", "Pamela", "Paraschiva", "Patricia", "Paula", "Paulica", "Paulina", "Petronela", "Petruta", "Pompilia", "Profira", "Rada", "Rafila", "Raluca", "Ramona", "Rebeca", "Renata", "Rica", "Roberta", "Robertina", "Rodica", "Romanita", "Romina", "Roxana", "Roxelana", "Roza", "Rozalia", "Ruxanda", "Ruxandra", "Sabina", "Sabrina", "Safta", "Salomea", "Sanda", "Saveta", "Savina", "Sânziana", "Semenica", "Severina", "Sidonia", "Silvana", "Silvia", "Silviana", "Simina", "Simona", "Smaranda", "Sofia", "Sonia", "Sorana", "Sorina", "Speranta", "Stana", "Stanca", "Stela", "Steliana", "Steluta", "Suzana", "Svetlana", "stefana", "stefania", "Tamara", "Tania", "Tatiana", "Teea", "Teodora", "Teodosia", "Teona", "Tiberia", "Timea", "Tinca", "Tincuta", "Tudora", "Tudorita", "Tudosia", "Valentina", "Valeria", "Vanesa", "Varvara", "Vasilica", "Venera", "Vera", "Veronica", "Veta", "Vicentia", "Victoria", "Violeta", "Viorela", "Viorica", "Virginia", "Viviana", "Vladelina", "Voichita", "Xenia", "Zaharia", "Zamfira", "Zaraza", "Zenobia", "Zenovia", "Zina", "Zoe", "Achim", "Adam", "Adelin", "Adi", "Adonis", "Adrian", "Agnos", "Albert", "Aleodor", "Alex", "Alexandru", "Alexe", "Alin", "Alistar", "Amedeu", "Amza", "Anatolie", "Andrei", "Andrian", "Angel", "Anghel", "Antim", "Anton", "Antonie", "Antoniu", "Arian", "Aristide", "Arsenie", "Augustin", "Aurel", "Aurelian", "Aurica", "Avram", "Axinte", "Barbu", "Bartolomeu", "Basarab", "Banel", "Bebe", "Beniamin", "Benone", "Bernard", "Bogdan", "Bradut", "Bucur", "Caius", "Camil", "Cantemir", "Carol", "Casian", "Cazimir", "Calin", "Catalin", "Cedrin", "Cezar", "Ciprian", "Claudiu", "Codin", "Codrin", "Codrut", "Constantin", "Cornel", "Corneliu", "Corvin", "Cosmin", "Costache", "Costel", "Costin", "Crin", "Cristea", "Cristian", "Cristobal", "Cristofor", "Dacian", "Damian", "Dan", "Daniel", "Darius", "David", "Decebal", "Denis", "Dinu", "Dominic", "Dorel", "Dorian", "Dorin", "Dorinel", "Doru", "Dragos", "Ducu", "Dumitru", "Edgar", "Edmond", "Eduard", "Eftimie", "Emanoil", "Emanuel", "Emanuil", "Emil", "Emilian", "Eremia", "Eric", "Ernest", "Eugen", "Eusebiu", "Eustatiu", "Fabian", "Felix", "Filip", "Fiodor", "Flaviu", "Florea", "Florentin", "Florian", "Florin", "Francisc", "Gabi", "Gabriel", "Gelu", "George", "Georgel", "Georgian", "Ghenadie", "Gheorghe", "Gheorghita", "Ghita", "Gica", "Gicu", "Giorgian", "Gratian", "Gregorian", "Grigore", "Haralamb", "Haralambie", "Horatiu", "Horea", "Horia", "Iacob", "Iancu", "Ianis", "Ieremia", "Ilarie", "Ilarion", "Ilie", "Inocentiu", "Ioan", "Ion", "Ionel", "Ionica", "Ionut", "Iosif", "Irinel", "Iulian", "Iuliu", "Iurie", "Iustin", "Iustinian", "Ivan", "Jan", "Jean", "Jenel", "Ladislau", "Lascar", "Laurentiu", "Laurian", "Lazar", "Leonard", "Leontin", "Leordean", "Lica", "Liviu", "Lorin", "Luca", "Lucentiu", "Lucian", "Lucretiu", "Ludovic", "Manole", "Marcel", "Marcu", "Marian", "Marin", "Marius", "Martin", "Matei", "Maxim", "Maximilian", "Madalin", "Mihai", "Mihail", "Mihnea", "Mina", "Mircea", "Miron", "Mitica", "Mitrut", "Mugur", "Mugurel", "Nae", "Narcis", "Nechifor", "Nelu", "Nichifor", "Nicoara", "Nicodim", "Nicolae", "Nicolaie", "Nicu", "Niculita", "Nicusor", "Nicuta", "Norbert", "Noris", "Norman", "Octav", "Octavian", "Octaviu", "Olimpian", "Olimpiu", "Oliviu", "Ovidiu", "Pamfil", "Panagachie", "Panait", "Paul", "Pavel", "Patru", "Petre", "Petrica", "Petrisor", "Petru", "Petrut", "Plesu", "Pompiliu", "Radu", "Rafael", "Rares", "Raul", "Raducu", "Razvan", "Relu", "Remus", "Robert", "Romeo", "Romulus", "Sabin", "Sandu", "Sandu", "Sava", "Sebastian", "Sergiu", "Sever", "Severin", "Silvian", "Silviu", "Simi", "Simion", "Sinica", "Sorin", "Stan", "Stancu", "Stelian", "serban", "stefan", "Teodor", "Teofil", "Teohari", "Theodor", "Tiberiu", "Timotei", "Titus", "Todor", "Toma", "Traian", "Tudor", "Valentin", "Valeriu", "Valter", "Vasile", "Vasilica", "Veniamin", "Vicentiu", "Victor", "Vincentiu", "Viorel", "Visarion", "Vlad", "Vladimir", "Vlaicu", "Voicu", "Zamfir", "Zeno"}
	lastNames := []string{"Aalboaiei", "Aalexoaiei", "Aamzar", "Aancutei", "Aandriesei", "Aanei", "Aanicai", "Aanii", "Aanitei", "Aanutei", "Aaron", "Aaxiniei", "Ababe", "Ababei", "Ababi", "Ababie", "Ababii", "Abadi", "Abadii", "Abageriu", "Abageru", "Abagiu", "Abahnencei", "Abaitancei", "Abaitanci", "Abalaie", "Abalaii", "Abalanesti", "Abalaoaie", "Abalaru", "Abalasei", "Abalasu", "Abaloaei", "Abaluta", "Abanaritei", "Abandei", "Abant", "Abatacesei", "Abaza", "Abdulea", "Abeaboeru", "Abeaboieru", "Abejenoaiei", "Abel", "Abela", "Abercei", "Abetegi", "Abetegii", "Abiculesei", "Abiculesi", "Abitanci", "Abitei", "Ablachim", "Ablai", "Aboaice", "Aboboaie", "Abobului", "Abogatoae", "Abogatoaie", "Abonculesei", "Abordeoaei", "Abordiencei", "Abordioaei", "Aborsoaie", "Abos", "Abosculesi", "Abostanoae", "Abouritei", "Abracel", "Abraham", "Abram", "Abramciuc", "Abramiuc", "Abramovici", "Abrasu", "Abriham", "Abrihan", "Abrii", "Abrudan", "Abrudanu", "Abrudean", "Abrudeana", "Abrudeanu", "Abrudian", "Abseleam", "Abucurei", "Abucuroaie", "Abudicioai", "Abuhaie", "Abuhnoaie", "Abuhnoaiei", "Abunei", "Abunoiu", "Abur", "Aburel", "Aburlacitei", "Abusan", "Abuseanu", "Abutnaritei", "Abutnariti", "Abutoaei", "Abutoaiei", "Abuzatoaei", "Abuzeloaei", "Acai", "Acalfoae", "Acalfoaie", "Acalinei", "Acambet", "Acamenitoaei", "Acapraritei", "Acarnaresei", "Acaroaie", "Acaru", "Acasandre", "Acasandrei", "Acasandrei", "Acasandri", "Acasandri", "Acasandrii", "Acatanoaie", "Acateu", "Acatiei", "Acatinca", "Acatinca", "Acatincai", "Acatincei", "Acatrinei", "Acatrini", "Acatrinii", "Acceleanu", "Accibas", "Acea", "Aceleanu", "Acelenescu", "Acerencu", "Acheaua", "Achelaritei", "Achetraritei", "Achiaitei", "Achiana", "Achifoaiei", "Achihaei", "Achihai", "Achihaie", "Achihaiei", "Achihaiei", "Achilaritei", "Achim", "Achimas", "Achimescu", "Achimet", "Achimia", "Achimut", "Achimuta", "Achinai", "Achinca", "Achiparoaiei", "Achirecesei", "Achirei", "Achiri", "Achiricesei", "Achiriloaei", "Achiriloaie", "Achiriloaiei", "Achiritoaei", "Achiroaei", "Achiroai", "Achiroaie", "Achiroaiei", "Achis", "Achitei", "Achitenei", "Achiti", "Acibotarita", "Acibotarita", "Acibotaritei", "Acichian", "Acim", "Acimovat", "Aciobanita", "Aciobanita", "Aciobanitei", "Aciobotaritei", "Aciocirlanoaiei", "Aciocoitei", "Acisoaie", "Aciu", "Aciuan", "Aciubotari", "Aciubotarita", "Aciubotaritei", "Acnim", "Acojocaritei", "Acolacioaie", "Acomi", "Acon", "Aconstantinese", "Aconstantinesei", "Aconstantinesi", "Aconutoae", "Aconutoaie", "Acornicesei", "Acosmei", "Acostache", "Acostachioae", "Acostachioaei", "Acostachioai", "Acostachioaie", "Acostachioaiei", "Acostantinesei", "Acosteoaie", "Acostinesei", "Acostioaei", "Acostoaei", "Acostoaie", "Acotunoaei", "Acotunoaie", "Acotunoaiei", "Acozmei", "Acreala", "Acretoaei", "Acriala", "Acris", "Acrismaritei", "Acrisor", "Acristei", "Acristinei", "Acroitorita", "Acroitorului", "Acru", "Acrudoae", "Acs", "Acsan", "Acsante", "Acsenciuc", "Acsenia", "Acsente", "Acsinciuc", "Acsinescu", "Acsinia", "Acsinia", "Acsinie", "Acsiniei", "Acsinii", "Acsinoiu", "Acsinte", "Acsintoaie", "Acu", "Acuculitei", "Acujboaei", "Acujboaie", "Acujboaiei", "Aculai", "Acunune", "Acununei", "Acununi", "Acurmoloae", "Adace", "Adafinei", "Adagiu", "Adalinean", "Adam", "Adamache", "Adamache", "Adamachi", "Adamcau", "Adamciuc", "Adamescu", "Adamescu", "Adamesteanu", "Adamet", "Adami", "Adamii", "Adamita", "Adamoae", "Adamoiu", "Adamovici", "Adamut", "Adamutiu", "Adan", "Adancu", "Adar", "Adascalita", "Adascalitei", "Adascaliti", "Adascalului", "Adavaloaie", "Adeaconitei", "Adelean", "Adelescu", "Adelman", "Aderca", "Adespei", "Adespii", "Adetu", "Adevar", "Adi", "Adiaconite", "Adiaconitei", "Adiaconiti", "Adiean", "Adiguzel", "Adimei", "Adincu", "Adir", "Adit", "Adjudeanu", "Admacof", "Adoamnei", "Adobritei", "Adobroaei", "Adobroaiei", "Adoc", "Adochiei", "Adochita", "Adochitei", "Adochiti", "Adogulesei", "Adoi", "Adominicai", "Adomnicai", "Adomnitei", "Adomniti", "Adomnitii", "Adonicioaie", "Adorean", "Adorian", "Adorjan", "Adorjani", "Adoroaei", "Adpredeanu", "Adragai", "Adragei", "Adrian", "Adriana", "Adronic", "Adroniciuc", "Adudec", "Aduducesei", "Aduloaei", "Aduloaia", "Aduloaie", "Adumitracesei", "Adumitrachioaie", "Adumitrachioaiei", "Adumitrei", "Adumitresei", "Adumitri", "Adumitroae", "Adumitroaei", "Adumitroaie", "Adumitroaiei", "Adurnoae", "Adusmanoaie", "Aecoboae", "Aecoboaei", "Aecoboaiei", "Aelenei", "Aeleni", "Aelisabetei", "Aenachioaie", "Aenculesei", "Aenoaei", "Aenoalei", "Aerimitoaia", "Aerinei", "Aeroaei", "Aeroaie", "Aevoae", "Aevoaei", "Afadaroaei", "Afanase", "Afanasie", "Afemeii", "Aferaritei", "Afetelor", "Afighiroaie", "Afilie", "Afilipoae", "Afilipoaei", "Afilipoaie", "Afilipoaiei", "Afiliu", "Afim", "Aflacailor", "Aflat", "Aflecailor", "Afloare", "Afloarei", "Afloari", "Afloarie", "Afloraritei", "Aflorei", "Afloria", "Afloroae", "Afloroaei", "Afloroaie", "Afloroaiei", "Afocsoaie", "Afodorcei", "Afoniu", "Afrapt", "Afrasaniei", "Afrasiloaia", "Afrasine", "Afrasinei", "Afrasinei", "Afrasini", "Afrem", "Afrentoaei", "Afrentoaie", "Afric", "Afrim", "Afronie", "Aftanache", "Aftanachi", "Aftanase", "Aftanase", "Aftanasie", "Aftene", "Afteni", "Aftenie", "Aftenii", "Afteri", "Afteu", "Aftin", "Aftinescu", "Aftion", "Aftode", "Aftodie", "Aftonia", "Aftonic", "Aftonie", "Afuduloae", "Afuduloai", "Afumateanu", "Afumatu", "Afusoae", "Afuza", "Aga", "Agache", "Agachi", "Agafiei", "Agafita", "Agafitei", "Agafiti", "Agaghe", "Againoaiei", "Agaleanu", "Aganencei", "Agape", "Agapescu", "Agapescu", "Agapi", "Agapia", "Agapie", "Agapin", "Agapsa", "Agarafinei", "Agarbicean", "Agarbiceanu", "Agarici", "Agarlita", "Agatinei", "Agavriloae", "Agavriloaei", "Agavriloai", "Agavriloaia", "Agavriloaie", "Agavriloaiei", "Ageroaei", "Ageroaie", "Ageroaiei", "Agescu", "Ageu", "Aghachi", "Aghapie", "Agheana", "Agheboaei", "Aghel", "Aghenitei", "Agheorgheoaiei", "Agheorghesei", "Agheorghiese", "Agheorghiesei", "Agheorghiesi", "Agheorghitoae", "Aghergheloaei", "Aghergheloaie", "Agherghinei", "Agherghinii", "Aghescu", "Aghiciuc", "Aghiculesei", "Aghiculesii", "Aghimescu", "Aghinei", "Aghiniei", "Aghinitei", "Aghioghiesei", "Aghiorghesei", "Aghiorghiesei", "Aghiorghioaei", "Aghiorghitoaie", "Aghiresan", "Aghis", "Aghitoaie", "Agiacai", "Agica", "Agigheoleanu", "Agighioleanu", "Agili", "Agimean", "Agiorghiesei", "Agiorgiuculesei", "Agirbicean", "Agirbiceanu", "Agiu", "Agiurgioaei", "Agliceru", "Aglici", "Aglitoiu", "Agopsa", "Agoroaei", "Agoroaie", "Agosoaie", "Agoston", "Agotici", "Agrapine", "Agrapine", "Agrapinei", "Agrapinei", "Agraviloaei", "Agreci", "Agres", "Agrici", "Agriesanu", "Agrigorcioaie", "Agrigoroae", "Agrigoroaei", "Agrigoroaie", "Agrigoroaiei", "Agrigoroei", "Agrijan", "Agrisan", "Agudaru", "Agudin", "Agulescu", "Agura", "Agurida", "Agus", "Agusoaei", "Ahanculesei", "Ahatinei", "Ahergheleghitei", "Aherghelegitiei", "Ahergheligita", "Ahergheligitei", "Aherghiligitei", "Ahorei", "Ahrisavoaei", "Ahritculese", "Ahritculesei", "Ahtamon", "Ahtemenciuc", "Ahuimanu", "Aiacoboae", "Aiacoboaei", "Aiacoboai", "Aiacoboaie", "Aiacoboaiei", "Aiacuboaei", "Aichimoaie", "Aicoboae", "Aida", "Aidanei", "Aidimireanu", "Aidoiu", "Aiecoboaie", "Aiecoboaiei", "Aiftimie", "Aiftimiei", "Aiftimitei", "Aiftimoae", "Aiftimoaei", "Aiftinca", "Aiftincai", "Aiftinica", "Aiftodoaie", "Aiftodoaiei", "Aignatoaie", "Aignatoaiei", "Ailene", "Ailenei", "Aileni", "Ailenii", "Ailiesa", "Ailiese", "Ailiesei", "Ailiesii", "Ailinca", "Ailincai", "Ailincutei", "Ailioaei", "Ailioai", "Ailioaie", "Ailioaiei", "Ailisoaie", "Ailoae", "Ailoaei", "Ailoaie", "Ailoaiei", "Ailutoaiei", "Ainoaiei", "Aioanei", "Aioani", "Aioanitoaie", "Aioji", "Aiojoaei", "Aiojoaiei", "Aioneasa", "Aionei", "Aionesa", "Aionesa", "Aionese", "Aionesei", "Aionicesei", "Aionitoaei", "Aionitoaie", "Aionoaie", "Aiordachiaei", "Aiordachioae", "Aiordachioaei", "Aiordachioai", "Aiordachioaie", "Aiovoae", "Airimioaei", "Airimioaie", "Airimioaiei", "Airimitoaie", "Airimitoaiei", "Airimoaei", "Airinei", "Airini", "Airoaei", "Airoaie", "Aitai", "Aitean", "Aitonean", "Aitoneanu", "Aiudeanu", "Aiuroaie", "Aiuseriu", "Aivancesei", "Aivanesa", "Aivanesei", "Aivanoaei", "Aivanoaie", "Aivanoaiei", "Ajarescu", "Ajder", "Ajidaucei", "Ajitaritai", "Ajitaritei", "Ajitariti", "Ajudeanu", "Ajudelui", "Alaci", "Alalitei", "Alamae", "Alaman", "Alamaru", "Alamiie", "Alamita", "Alamorean", "Alamoreanu", "Alamureanu", "Alazaroae", "Alazaroaei", "Alazaroaie", "Alazaroei", "Alb", "Alba", "Albac", "Albaceanu", "Albai", "Albali", "Alban", "Albanas", "Albani", "Albastrel", "Albastroiu", "Albea", "Albean", "Albeanu", "Albei", "Albeiu", "Albertus", "Albescu", "Albescu ", "Albesteanu", "Albetel", "Albi", "Albici", "Albiciu", "Albieru", "Albina", "Albinaru", "Albinet", "Albis", "Albisi", "Albisor", "Albisoru", "Albisteanu", "Albita", "Alboaie", "Alboanu", "Alboi", "Alboiu", "Albot", "Albota", "Albotoaei", "Albu", "Albuica", "Albuleanu", "Albulescu", "Albulet", "Albulete", "Albuletu", "Albureanu", "Albus", "Albut", "Albuta", "Albutiu", "Alcaziu", "Alcea", "Alda", "Aldea", "Aldecu", "Aldescu", "Aldica", "Aldis", "Aldoi", "Aldoiu", "Aldulea", "Aldulescu", "Aleca", "Alecsa", "Alecsandrescu", "Alecsandrescu", "Alecsandri", "Alecsandriescu", "Alecsandroaie", "Alecsandru", "Alecse", "Alecseiciuc", "Alecsoae", "Alecsoaei", "Alecsoai", "Alecsoaia", "Alecsoaie", "Alecsoiu", "Alecsuc", "Alecsuta", "Alecu", "Alecus", "Alecusan", "Aleman", "Alemnaritei", "Alenei", "Aleonte", "Aleosan", "Alergus", "Alesandru", "Alessandrescu", "Alesteu", "Alesu", "Alesutan", "Alexa", "Alexa", "Alexai", "Alexan", "Alexana", "Alexandra", "Alexandrache"}

	randName := RandName{
		FirstName: firstNames[rand.Intn(len(firstNames))],
		LastName:  lastNames[rand.Intn(len(lastNames))],
	}

	return randName
}
