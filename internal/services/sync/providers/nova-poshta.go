package providers

import (
	"sashapekh/country_delivery_service/internal/repositories"
	"sashapekh/country_delivery_service/pkg/novaposhta"
	"time"

	"github.com/gosimple/slug"
)

func SyncRegionsNovaposhta(
	np *novaposhta.Novaposhta,
	repoH *repositories.RepoHandler,
) error {
	regions, err := np.GetAllRegions()

	if err != nil {
		return err
	}

	for _, region := range regions {

		mongoRegion := repositories.Region{
			Name:      region.Description,
			Slug:      slug.Make(region.Description),
			NpRef:     region.Ref,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		}

		repoH.InsertRegion(mongoRegion)
	}

	return nil
}
