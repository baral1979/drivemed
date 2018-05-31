package main

type productFr struct {
	ProductCodeSKU                    string `json:"productcodesku"`
	FamilyIdentifier                  string `json:"familyidentifier"`
	FamilyProductTitle                string `json:"familyproducttitle"`
	IndividualProductTitle            string `json:"individualproducttitle"`
	RootCategory                      string `json:"rootcategory"`
	ChildCategory                     string `json:"childcategory"`
	GrandchildCategory                string `json:"grandchildcategory"`
	Category                          string `json:"category"`
	ProductUPC                        string `json:"productupc"`
	GTINITFBarcode                    string `json:"gtinitfbarcode"`
	ConsumerUnitofMeasure             string `json:"consumerunitofmeasure"`
	MAPMinimumAdvertisedPrice         string `json:"mapminimumadvertisedprice"`
	Taxation                          string `json:"taxation"`
	IndividualProductDescription      string `json:"individualproductdescription"`
	IndividualProductFeatures         string `json:"individualproductfeatures"`
	FamilyProductAttributes           string `json:"familyproductattributes"`
	FamilyProductDescription          string `json:"familyproductdescription"`
	FamilyProductFeatures             string `json:"familyproductfeatures"`
	ProductWarrantyInfo               string `json:"productwarrantyinfo"`
	CountryofOrigin                   string `json:"countryoforigin"`
	ProductAssembly                   string `json:"productassembly"`
	ProductInstallation               string `json:"productinstallation"`
	Proposition65Warning              string `json:"proposition65warning"`
	RequiresPrescription              string `json:"requiresprescription"`
	HygienicProduct                   string `json:"hygienicproduct"`
	RestockingFee                     string `json:"restockingfee"`
	GoingGreen                        string `json:"goinggreen"`
	Allergy                           string `json:"allergy"`
	ShippingLeadTime                  string `json:"shippingleadtime"`
	PrimaryShippingMethod             string `json:"primaryshippingmethod"`
	NumberofShippingBoxes             string `json:"numberofshippingboxes"`
	PrimaryShippingBoxLength          string `json:"primaryshippingboxlength"`
	PrimaryShippingBoxWidth           string `json:"primaryshippingboxwidth"`
	PrimaryShippingBoxHeight          string `json:"primaryshippingboxheight"`
	PrimaryShippingBoxWeight          string `json:"primaryshippingboxweight"`
	ShippingBoxLength2                string `json:"shippingboxlength2"`
	ShippingBoxWidth2                 string `json:"shippingboxwidth2"`
	ShippingBoxHeight2                string `json:"shippingboxheight2"`
	ShippingBoxWeight2                string `json:"shippingboxweight2"`
	ShippingBoxLength3                string `json:"shippingboxlength3"`
	ShippingBoxWidth3                 string `json:"shippingboxwidth3"`
	ShippingBoxHeight3                string `json:"shippingboxheight3"`
	ShippingBoxWeight3                string `json:"shippingboxweight3"`
	ShippingBoxLength4                string `json:"shippingboxlength4"`
	ShippingBoxWidth4                 string `json:"shippingboxwidth4"`
	ShippingBoxHeight4                string `json:"shippingboxheight4"`
	ShippingBoxWeight4                string `json:"shippingboxweight4"`
	ShippingBoxLength5                string `json:"shippingboxlength5"`
	ShippingBoxWidth5                 string `json:"shippingboxwidth5"`
	ShippingBoxHeight5                string `json:"shippingboxheight5"`
	ShippingBoxWeight5                string `json:"shippingboxweight5"`
	ShippingBoxLength6                string `json:"shippingboxlength6"`
	ShippingBoxWidth6                 string `json:"shippingboxwidth6"`
	ShippingBoxHeight6                string `json:"shippingboxheight6"`
	ShippingBoxWeight6                string `json:"shippingboxweight6"`
	ShippingBoxLength7                string `json:"shippingboxlength7"`
	ShippingBoxWidth7                 string `json:"shippingboxwidth7"`
	ShippingBoxHeight7                string `json:"shippingboxheight7"`
	ShippingBoxWeight7                string `json:"shippingboxweight7"`
	ShippingBoxLength8                string `json:"shippingboxlength8"`
	ShippingBoxWidth8                 string `json:"shippingboxwidth8"`
	ShippingBoxHeight8                string `json:"shippingboxheight8"`
	ShippingBoxWeight8                string `json:"shippingboxweight8"`
	ShippingBoxLength9                string `json:"shippingboxlength9"`
	ShippingBoxWidth9                 string `json:"shippingboxwidth9"`
	ShippingBoxHeight9                string `json:"shippingboxheight9"`
	ShippingBoxWeight9                string `json:"shippingboxweight9"`
	ShippingBoxLength10               string `json:"shippingboxlength10"`
	ShippingBoxWidth10                string `json:"shippingboxwidth10"`
	ShippingBoxHeight10               string `json:"shippingboxheight10"`
	ShippingBoxWeight10               string `json:"shippingboxweight10"`
	ManufacturerName                  string `json:"manufacturername"`
	ManufacturerDescription           string `json:"manufacturerdescription"`
	MissionStatement                  string `json:"missionstatement"`
	BrandDescription                  string `json:"branddescription"`
	OverallProductLength              string `json:"overallproductlength"`
	OverallProductWidth               string `json:"overallproductwidth"`
	OverallProductHeight              string `json:"overallproductheight"`
	ActualProductWeight               string `json:"actualproductweight"`
	PrimaryProductColor               string `json:"primaryproductcolor"`
	PrimaryProductMaterial            string `json:"primaryproductmaterial"`
	ProductWeightCapacity             string `json:"productweightcapacity"`
	ClosedWidth                       string `json:"closedwidth"`
	SeatWidth                         string `json:"seatwidth"`
	SeatDepth                         string `json:"seatdepth"`
	SeattoFloorHeight                 string `json:"seattofloorheight"`
	BackofChairHeight                 string `json:"backofchairheight"`
	NumberofWheels                    string `json:"numberofwheels"`
	WheelchairWeightWithoutBattery    string `json:"wheelchairweightwithoutbattery"`
	SteeringType                      string `json:"steeringtype"`
	MaxSpeed                          string `json:"maxspeed"`
	WidthInsideBackLegs               string `json:"widthinsidebacklegs"`
	InsideHandGripWidth               string `json:"insidehandgripwidth"`
	OverallLengthReclined             string `json:"overalllengthreclined"`
	AdjustableClamp                   string `json:"adjustableclamp"`
	AdjustableHeight                  string `json:"adjustableheight"`
	HeightAdjustments                 string `json:"heightadjustments"`
	AdjustableLength                  string `json:"adjustablelength"`
	AntiTipWheels                     string `json:"antitipwheels"`
	SeattoArmrestHeight               string `json:"seattoarmrestheight"`
	ArmrestLength                     string `json:"armrestlength"`
	ArmresttoFloorHeight              string `json:"armresttofloorheight"`
	WidthBetweenArms                  string `json:"widthbetweenarms"`
	RearWheels                        string `json:"rearwheels"`
	BackofChairWidth                  string `json:"backofchairwidth"`
	BackrestReclinesto                string `json:"backrestreclinesto"`
	BagHeight                         string `json:"bagheight"`
	BagLength                         string `json:"baglength"`
	BagMaterial                       string `json:"bagmaterial"`
	BagWidth                          string `json:"bagwidth"`
	BarLength                         string `json:"barlength"`
	BaseDepth                         string `json:"basedepth"`
	BaseHeight                        string `json:"baseheight"`
	BaseWidth                         string `json:"basewidth"`
	BaseClearance                     string `json:"baseclearance"`
	BaseWidthClosed                   string `json:"basewidthclosed"`
	BatteryDuration                   string `json:"batteryduration"`
	Batteries                         string `json:"batteries"`
	BatteryWeight                     string `json:"batteryweight"`
	BaseWeight                        string `json:"baseweight"`
	FoldedBasketHeight                string `json:"foldedbasketheight"`
	FoldedBasketLength                string `json:"foldedbasketlength"`
	FoldedBasketWidth                 string `json:"foldedbasketwidth"`
	BasketWidth                       string `json:"basketwidth"`
	BasketHeight                      string `json:"basketheight"`
	BasketLength                      string `json:"basketlength"`
	BedHeightMin                      string `json:"bedheightmin"`
	BedHeightMax                      string `json:"bedheightmax"`
	BoardLength                       string `json:"boardlength"`
	BoardWidth                        string `json:"boardwidth"`
	BoardHeight                       string `json:"boardheight"`
	BoomHeight                        string `json:"boomheight"`
	BoomLength                        string `json:"boomlength"`
	Brakes                            string `json:"brakes"`
	HandleHeightMin                   string `json:"handleheightmin"`
	HandleHeightMax                   string `json:"handleheightmax"`
	HandleLength                      string `json:"handlelength"`
	HandlebartoArmrestHeight          string `json:"handlebartoarmrestheight"`
	Volume                            string `json:"volume"`
	Case                              string `json:"case"`
	Casters                           string `json:"casters"`
	CatheterTube                      string `json:"cathetertube"`
	MachineTube                       string `json:"machinetube"`
	FitsCeilingHeights                string `json:"fitsceilingheights"`
	Charger                           string `json:"charger"`
	ChargingTime                      string `json:"chargingtime"`
	ClimbingAngle                     string `json:"climbingangle"`
	ContractionTime                   string `json:"contractiontime"`
	Controller                        string `json:"controller"`
	CoverMaterial                     string `json:"covermaterial"`
	CradlePoints                      string `json:"cradlepoints"`
	MaximumRange                      string `json:"maximumrange"`
	DedicatedHeelSection              string `json:"dedicatedheelsection"`
	OutsideLegsDepth                  string `json:"outsidelegsdepth"`
	OutsideLegsWidth                  string `json:"outsidelegswidth"`
	Design                            string `json:"design"`
	PumpLength                        string `json:"pumplength"`
	PumpHeight                        string `json:"pumpheight"`
	PumpWidth                         string `json:"pumpwidth"`
	Diameter                          string `json:"diameter"`
	FoldedDimensions                  string `json:"foldeddimensions"`
	PulseRateMeasurementRange         string `json:"pulseratemeasurementrange"`
	PulseRateAccuracyRange            string `json:"pulserateaccuracyrange"`
	SpO2MeasurementRange              string `json:"spo2measurementrange"`
	SpO2AccuracyRange                 string `json:"spo2accuracyrange"`
	PulseRateDisplayRange             string `json:"pulseratedisplayrange"`
	DisplayScreenLength               string `json:"displayscreenlength"`
	DisplayScreenWidth                string `json:"displayscreenwidth"`
	DisplayType                       string `json:"displaytype"`
	LEDDisplayColor                   string `json:"leddisplaycolor"`
	DistanceBetweenHipGuides          string `json:"distancebetweenhipguides"`
	DistanceBetweenLaterals           string `json:"distancebetweenlaterals"`
	Motor                             string `json:"motor"`
	DriveWheels                       string `json:"drivewheels"`
	ElectricalRequirements            string `json:"electricalrequirements"`
	EnergySaver                       string `json:"energysaver"`
	FilterType                        string `json:"filtertype"`
	FingertipWidthRange               string `json:"fingertipwidthrange"`
	FireRating                        string `json:"firerating"`
	FitsCylinders                     string `json:"fitscylinders"`
	FrontWheels                       string `json:"frontwheels"`
	GelType                           string `json:"geltype"`
	GroundClearance                   string `json:"groundclearance"`
	HarnessDepth                      string `json:"harnessdepth"`
	HarnessWidth                      string `json:"harnesswidth"`
	HeaviestPiece                     string `json:"heaviestpiece"`
	HipGuidesDepth                    string `json:"hipguidesdepth"`
	HipGuidesHeight                   string `json:"hipguidesheight"`
	HoseLength                        string `json:"hoselength"`
	OperatingRelativeHumidity         string `json:"operatingrelativehumidity"`
	InletPressure                     string `json:"inletpressure"`
	IntensityControl                  string `json:"intensitycontrol"`
	LateralDepth                      string `json:"lateraldepth"`
	LateralHeight                     string `json:"lateralheight"`
	LayersofFoam                      string `json:"layersoffoam"`
	OverallLengthwRiggings            string `json:"overalllengthwriggings"`
	LiterFlow                         string `json:"literflow"`
	LiterFlowIncrements               string `json:"literflowincrements"`
	PumpAirflow                       string `json:"pumpairflow"`
	LowPowerConsumption               string `json:"lowpowerconsumption"`
	MaximumPressure                   string `json:"maximumpressure"`
	MethodofNebulization              string `json:"methodofnebulization"`
	NeckStrapLength                   string `json:"neckstraplength"`
	NumberofSections                  string `json:"numberofsections"`
	OperatingAngle                    string `json:"operatingangle"`
	OperatingPressure                 string `json:"operatingpressure"`
	OperatingTemperature              string `json:"operatingtemperature"`
	OutletType                        string `json:"outlettype"`
	OutletPressure                    string `json:"outletpressure"`
	SeattoFootDeck                    string `json:"seattofootdeck"`
	OxygenPressure                    string `json:"oxygenpressure"`
	PadWidth                          string `json:"padwidth"`
	PadHeight                         string `json:"padheight"`
	PadLength                         string `json:"padlength"`
	ParticleSize                      string `json:"particlesize"`
	PowerRequirements                 string `json:"powerrequirements"`
	ProductSize                       string `json:"productsize"`
	ProtectiveBootMaterial            string `json:"protectivebootmaterial"`
	PulseAmplitude                    string `json:"pulseamplitude"`
	PulseFrequency                    string `json:"pulsefrequency"`
	PulseWidth                        string `json:"pulsewidth"`
	PumpAlarms                        string `json:"pumpalarms"`
	PumpCycleTime                     string `json:"pumpcycletime"`
	PumpPower                         string `json:"pumppower"`
	RailDepth                         string `json:"raildepth"`
	RailHeight                        string `json:"railheight"`
	RailWidth                         string `json:"railwidth"`
	RampTime                          string `json:"ramptime"`
	RecommendedUserHeight             string `json:"recommendeduserheight"`
	RelaxationTime                    string `json:"relaxationtime"`
	PumpType                          string `json:"pumptype"`
	FloortoFootrest                   string `json:"floortofootrest"`
	SeattoLaterals                    string `json:"seattolaterals"`
	SeatWeight                        string `json:"seatweight"`
	Shape                             string `json:"shape"`
	SideRails                         string `json:"siderails"`
	SlingPoints                       string `json:"slingpoints"`
	StepHeight                        string `json:"stepheight"`
	StepSurfaceDepth                  string `json:"stepsurfacedepth"`
	StepSurfaceWidth                  string `json:"stepsurfacewidth"`
	StorageTemperature                string `json:"storagetemperature"`
	TabletopWidth                     string `json:"tabletopwidth"`
	TabletopDepth                     string `json:"tabletopdepth"`
	TabletopHeight                    string `json:"tabletopheight"`
	TurningRadius                     string `json:"turningradius"`
	UnderarmHeight                    string `json:"underarmheight"`
	UnitHeight                        string `json:"unitheight"`
	UnitHeightwithProtectiveBootWidth string `json:"unitheightwithprotectivebootwidth"`
	UnitLength                        string `json:"unitlength"`
	UnitLengthwithProtectiveBootWidth string `json:"unitlengthwithprotectivebootwidth"`
	UnitWidth                         string `json:"unitwidth"`
	UnitWidthwithProtectiveBootWidth  string `json:"unitwidthwithprotectivebootwidth"`
	FreewheelMode                     string `json:"freewheelmode"`
	BedHeightwCastersMin              string `json:"bedheightwcastersmin"`
	BedHeightwCastersMax              string `json:"bedheightwcastersmax"`
	GroundClearancewCasters           string `json:"groundclearancewcasters"`
	HeadboardHeightMin                string `json:"headboardheightmin"`
	HeadboardHeightwCastersMin        string `json:"headboardheightwcastersmin"`
	HeadboardHeightMax                string `json:"headboardheightmax"`
	HeadboardHeightwCastersMax        string `json:"headboardheightwcastersmax"`
	FootboardHeightMin                string `json:"footboardheightmin"`
	FootboardHeightwCastersMin        string `json:"footboardheightwcastersmin"`
	FootboardHeightMax                string `json:"footboardheightmax"`
	FootboardHeightwCastersMax        string `json:"footboardheightwcastersmax"`
	HeadSectionHeightMin              string `json:"headsectionheightmin"`
	HeadSectionHeightwCastersMin      string `json:"headsectionheightwcastersmin"`
	HeadSectionHeightMax              string `json:"headsectionheightmax"`
	HeadSectionHeightwCastersMax      string `json:"headsectionheightwcastersmax"`
	FootSectionHeightMin              string `json:"footsectionheightmin"`
	FootSectionHeightwCastersMin      string `json:"footsectionheightwcastersmin"`
	FootSectionHeightMax              string `json:"footsectionheightmax"`
	FootSectionHeightwCastersMax      string `json:"footsectionheightwcastersmax"`
	WheelchairWeightWithoutRiggings   string `json:"wheelchairweightwithoutriggings"`
	OscillationRate                   string `json:"oscillationrate"`
	WidthBetweenArmrestPads           string `json:"widthbetweenarmrestpads"`
	WidthBetweenPosts                 string `json:"widthbetweenposts"`
	WidthofSeatUpholstery             string `json:"widthofseatupholstery"`
	DepthofSeatUpholstery             string `json:"depthofseatupholstery"`
	MoistureOutput                    string `json:"moistureoutput"`
	RecommendedRoomSize               string `json:"recommendedroomsize"`
	FitsMattressSizes                 string `json:"fitsmattresssizes"`
	InsideTrackWidth                  string `json:"insidetrackwidth"`
}
