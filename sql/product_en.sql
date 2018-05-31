select
    ProductCodeSKU AS 'Product Code',
    ProductUPC AS 'Product UPC',
    FamilyIdentifier AS 'Product Family',
    IndividualProductTitle AS 'Product Name',
    IndividualProductDescription AS 'Product Description',
    IndividualProductFeatures AS 'Product Features',
    FamilyProductFeatures AS 'Product Options',
    CASE WHEN RootCategory = '' THEN 'Other' ELSE RootCategory END AS 'Root Category',
    ChildCategory AS 'Child Category',
    0 AS 'Price', --todo
    Taxation AS 'Taxable',
    PrimaryImageFileName AS 'Product Image File Name',
    AdditionalImageFileName1 AS 'Additional Image File Name 1',
    AdditionalImageFileName2 AS 'Additional Image File Name 2',
    AdditionalImageFileName3 AS 'Additional Image File Name 3',
    AdditionalImageFileName4 AS 'Additional Image File Name 4',
    AdditionalImageFileName5 AS 'Additional Image File Name 5',
    AdditionalImageFileName6 AS 'Additional Image File Name 6',
    AdditionalImageFileName7 AS 'Additional Image File Name 7',
    AdditionalImageFileName8 AS 'Additional Image File Name 8',
    AdditionalImageFileName9 AS 'Additional Image File Name 9',
    PrimaryYouTubeInformationalVideoFileLink AS 'Product Video File Name',
    BrandName AS 'Brand Name',
    ProductWarrantyInfo AS 'Warranty',
    RequiresPrescription AS 'Requires Prescription',
    CASE WHEN PrimaryImageFileName = '' THEN '' ELSE 'thumb_' + PrimaryImageFileName  END AS 'Product Thumbnail File Name',
    Allergy AS 'Allergy',
    0 AS 'Cost', --todo
    0 AS 'Shipping Cost', --todo
    'No' AS 'Missing Translation', --todo
    0 AS 'Product Significance' --todo
from
    SourceEn