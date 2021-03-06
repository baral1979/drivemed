package main

type productEn struct {
	ProductCodeSKU                              string `json:"productcodesku"`
	FamilyIdentifier                            string `json:"familyidentifier"`
	FamilyProductTitle                          string `json:"familyproducttitle"`
	IndividualProductTitle                      string `json:"individualproducttitle"`
	RootCategory                                string `json:"rootcategory"`
	ChildCategory                               string `json:"childcategory"`
	GrandchildCategory                          string `json:"grandchildcategory"`
	Category                                    string `json:"category"`
	ProductUPC                                  string `json:"productupc"`
	GTINITFBarcode                              string `json:"gtinitfbarcode"`
	UnitofMeasure                               string `json:"unitofmeasure"`
	MAPMinimumAdvertisedPrice                   string `json:"mapminimumadvertisedprice"`
	MSRPSuggestedRetailPrice                    string `json:"msrpsuggestedretailprice"`
	Taxation                                    string `json:"taxation"`
	CombinedFamilyHTML                          string `json:"combinedfamilyhtml"`
	CombinedIndividualHTML                      string `json:"combinedindividualhtml"`
	IndividualProductDescription                string `json:"individualproductdescription"`
	IndividualProductFeatures                   string `json:"individualproductfeatures"`
	IndividualProductSpecs                      string `json:"individualproductspecs"`
	FamilyProductAttributes                     string `json:"familyproductattributes"`
	FamilyProductDescription                    string `json:"familyproductdescription"`
	FamilyProductFeatures                       string `json:"familyproductfeatures"`
	FamilyProductSpecs                          string `json:"familyproductspecs"`
	OverallProductLength                        string `json:"overallproductlength"`
	OverallProductWidth                         string `json:"overallproductwidth"`
	OverallProductHeight                        string `json:"overallproductheight"`
	ActualProductWeight                         string `json:"actualproductweight"`
	ProductWeightCapacity                       string `json:"productweightcapacity"`
	ProductWarrantyInfo                         string `json:"productwarrantyinfo"`
	CountryofOrigin                             string `json:"countryoforigin"`
	PrimaryProductMaterial                      string `json:"primaryproductmaterial"`
	PrimaryProductColor                         string `json:"primaryproductcolor"`
	HCPCSBillingCode                            string `json:"hcpcsbillingcode"`
	FSAEligible                                 string `json:"fsaeligible"`
	UserSize                                    string `json:"usersize"`
	ProductAssembly                             string `json:"productassembly"`
	ProductInstallation                         string `json:"productinstallation"`
	Proposition65Warning                        string `json:"proposition65warning"`
	RequiresPrescription                        string `json:"requiresprescription"`
	HygienicProduct                             string `json:"hygienicproduct"`
	RestockingFee                               string `json:"restockingfee"`
	GoingGreen                                  string `json:"goinggreen"`
	Allergy                                     string `json:"allergy"`
	Keywords                                    string `json:"keywords"`
	CrossSelling                                string `json:"crossselling"`
	PrimaryImageFileName                        string `json:"primaryimagefilename"`
	PrimaryImageLink                            string `json:"primaryimagelink"`
	AdditionalImageFileName1                    string `json:"additionalimagefilename1"`
	SubImageLink1                               string `json:"subimagelink1"`
	AdditionalImageFileName2                    string `json:"additionalimagefilename2"`
	SubImageLink2                               string `json:"subimagelink2"`
	AdditionalImageFileName3                    string `json:"additionalimagefilename3"`
	SubImageLink3                               string `json:"subimagelink3"`
	AdditionalImageFileName4                    string `json:"additionalimagefilename4"`
	SubImageLink4                               string `json:"subimagelink4"`
	AdditionalImageFileName5                    string `json:"additionalimagefilename5"`
	SubImageLink5                               string `json:"subimagelink5"`
	AdditionalImageFileName6                    string `json:"additionalimagefilename6"`
	SubImageLink6                               string `json:"subimagelink6"`
	AdditionalImageFileName7                    string `json:"additionalimagefilename7"`
	SubImageLink7                               string `json:"subimagelink7"`
	AdditionalImageFileName8                    string `json:"additionalimagefilename8"`
	SubImageLink8                               string `json:"subimagelink8"`
	AdditionalImageFileName9                    string `json:"additionalimagefilename9"`
	SubImageLink9                               string `json:"subimagelink9"`
	AdditionalImageFileName10                   string `json:"additionalimagefilename10"`
	SubImageLink10                              string `json:"subimagelink10"`
	AdditionalImageFileName11                   string `json:"additionalimagefilename11"`
	SubImageLink11                              string `json:"subimagelink11"`
	AdditionalImageFileName12                   string `json:"additionalimagefilename12"`
	SubImageLink12                              string `json:"subimagelink12"`
	AdditionalImageFileName13                   string `json:"additionalimagefilename13"`
	SubImageLink13                              string `json:"subimagelink13"`
	AdditionalImageFileName14                   string `json:"additionalimagefilename14"`
	SubImageLink14                              string `json:"subimagelink14"`
	AdditionalImageFileName15                   string `json:"additionalimagefilename15"`
	SubImageLink15                              string `json:"subimagelink15"`
	AdditionalImageFileName16                   string `json:"additionalimagefilename16"`
	SubImageLink16                              string `json:"subimagelink16"`
	AdditionalImageFileName17                   string `json:"additionalimagefilename17"`
	SubImageLink17                              string `json:"subimagelink17"`
	AdditionalImageFileName18                   string `json:"additionalimagefilename18"`
	SubImageLink18                              string `json:"subimagelink18"`
	AdditionalImageFileName19                   string `json:"additionalimagefilename19"`
	SubImageLink19                              string `json:"subimagelink19"`
	ProductManualFileName                       string `json:"productmanualfilename"`
	ProductManualLink                           string `json:"productmanuallink"`
	RotationalImageLink                         string `json:"rotationalimagelink"`
	PrimaryInformationalVideoFileName           string `json:"primaryinformationalvideofilename"`
	PrimaryDriveInformationalVideoFileLink      string `json:"primarydriveinformationalvideofilelink"`
	PrimaryYouTubeInformationalVideoFileLink    string `json:"primaryyoutubeinformationalvideofilelink"`
	SecondaryDriveInformationalVideoFileName    string `json:"secondarydriveinformationalvideofilename"`
	SecondaryDriveInformationalVideoFileLink    string `json:"secondarydriveinformationalvideofilelink"`
	SecondaryYouTubeInformationalVideoFileLink  string `json:"secondaryyoutubeinformationalvideofilelink"`
	AssemblyVideoFileName                       string `json:"assemblyvideofilename"`
	AssemblyVideoFileLink                       string `json:"assemblyvideofilelink"`
	ShippingLeadTime                            string `json:"shippingleadtime"`
	PrimaryShippingMethod                       string `json:"primaryshippingmethod"`
	NumberofShippingBoxes                       string `json:"numberofshippingboxes"`
	PrimaryShippingBoxLength                    string `json:"primaryshippingboxlength"`
	PrimaryShippingBoxWidth                     string `json:"primaryshippingboxwidth"`
	PrimaryShippingBoxHeight                    string `json:"primaryshippingboxheight"`
	PrimaryShippingBoxWeight                    string `json:"primaryshippingboxweight"`
	ShippingBoxLength2                          string `json:"shippingboxlength2"`
	ShippingBoxWidth2                           string `json:"shippingboxwidth2"`
	ShippingBoxHeight2                          string `json:"shippingboxheight2"`
	ShippingBoxWeight2                          string `json:"shippingboxweight2"`
	ShippingBoxLength3                          string `json:"shippingboxlength3"`
	ShippingBoxWidth3                           string `json:"shippingboxwidth3"`
	ShippingBoxHeight3                          string `json:"shippingboxheight3"`
	ShippingBoxWeight3                          string `json:"shippingboxweight3"`
	ShippingBoxLength4                          string `json:"shippingboxlength4"`
	ShippingBoxWidth4                           string `json:"shippingboxwidth4"`
	ShippingBoxHeight4                          string `json:"shippingboxheight4"`
	ShippingBoxWeight4                          string `json:"shippingboxweight4"`
	ShippingBoxLength5                          string `json:"shippingboxlength5"`
	ShippingBoxWidth5                           string `json:"shippingboxwidth5"`
	ShippingBoxHeight5                          string `json:"shippingboxheight5"`
	ShippingBoxWeight5                          string `json:"shippingboxweight5"`
	ShippingBoxLength6                          string `json:"shippingboxlength6"`
	ShippingBoxWidth6                           string `json:"shippingboxwidth6"`
	ShippingBoxHeight6                          string `json:"shippingboxheight6"`
	ShippingBoxWeight6                          string `json:"shippingboxweight6"`
	ShippingBoxLength7                          string `json:"shippingboxlength7"`
	ShippingBoxWidth7                           string `json:"shippingboxwidth7"`
	ShippingBoxHeight7                          string `json:"shippingboxheight7"`
	ShippingBoxWeight7                          string `json:"shippingboxweight7"`
	ShippingBoxLength8                          string `json:"shippingboxlength8"`
	ShippingBoxWidth8                           string `json:"shippingboxwidth8"`
	ShippingBoxHeight8                          string `json:"shippingboxheight8"`
	ShippingBoxWeight8                          string `json:"shippingboxweight8"`
	ShippingBoxLength9                          string `json:"shippingboxlength9"`
	ShippingBoxWidth9                           string `json:"shippingboxwidth9"`
	ShippingBoxHeight9                          string `json:"shippingboxheight9"`
	ShippingBoxWeight9                          string `json:"shippingboxweight9"`
	ShippingBoxLength10                         string `json:"shippingboxlength10"`
	ShippingBoxWidth10                          string `json:"shippingboxwidth10"`
	ShippingBoxHeight10                         string `json:"shippingboxheight10"`
	ShippingBoxWeight10                         string `json:"shippingboxweight10"`
	ManufacturerName                            string `json:"manufacturername"`
	ManufacturerDescription                     string `json:"manufacturerdescription"`
	MissionStatement                            string `json:"missionstatement"`
	ManufacturerLogoImage                       string `json:"manufacturerlogoimage"`
	ManufacturerLogoImageLink                   string `json:"manufacturerlogoimagelink"`
	ManufacturerDescriptionVideoFileName        string `json:"manufacturerdescriptionvideofilename"`
	ManufacturerDescriptionDriveVideoFileLink   string `json:"manufacturerdescriptiondrivevideofilelink"`
	ManufacturerDescriptionYouTubeVideoFileLink string `json:"manufacturerdescriptionyoutubevideofilelink"`
	BrandName                                   string `json:"brandname"`
	BrandDescription                            string `json:"branddescription"`
	BrandLogoImage                              string `json:"brandlogoimage"`
	BrandlogoImageLink                          string `json:"brandlogoimagelink"`
}
